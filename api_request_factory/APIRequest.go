package api_request_factory

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/BrandonEchols/common-go-utils/common"
	common_routing "github.com/BrandonEchols/common-go-utils/routing"
	"io/ioutil"
	"net/http"
	"time"
)

type IAPIRequest interface {
	//This is an interface so that we can mock it out and not actually 'Do' it in tests
	DoAsync(common.IResult)
	Do() common.IResult
	GetHttpResponse() *http.Response
}

//An APIRequest struct is meant to be an all-inclusive object that sends/handles API requests from server to server.
type APIRequest struct {
	/*
		Request and util fields
	*/
	//The IConfigGetter to use
	config common.IConfigGetter
	//The http.Client to use
	client *http.Client
	//The Url to make the request to
	Url string
	//The http Method of the request
	Method string
	//The request body to json.Marshal and send, or nil
	RequestBody interface{}
	//The IRequestFormatter to apply to the request when 'Do'ing
	Formatter common.IRequestFormatter
	//The Request Headers
	Headers map[string]string
	//The Request Context
	Context context.Context
	//The Semantic name of the API the request is calling (used for contextual logging/error response message)
	ApiName string
	/*
		Validation and response fields
	*/
	//A map of http.status codes to ExpectedResponseBody's or nil. If the ExpectedResponseBody is an implementation of an
	//IPayload, it will validate when 'do'ing the request
	ValidResponses map[int]interface{}
	//The http.Response to the request
	HttpResponse *http.Response
	//The number of times the request should be retried if the return body is invalid or the response code is > 499
	NumTries int
	//The number of milliseconds to wait inbetween retries
	DelayBetweenTries int
	//The number of characters of the response body to log
	ResponseLogLimit int
}

type IPayload interface {
	Valid() error
}

/*
	getBaseAPIRequest Builds a standard APIRequest with the needed injection and defaults.
	@params
		client *http.Client The http.Client to use to make the requests
		config common.IConfigGetter The configGetter to inject into the request
		url string The full url to send the request to
		method string The http Method to send the request as
*/
func getBaseAPIRequest(client *http.Client, config common.IConfigGetter, url string, method string) *APIRequest {
	return &APIRequest{
		client:            client,
		config:            config,
		Url:               url,
		Method:            method,
		Headers:           map[string]string{},
		ValidResponses:    map[int]interface{}{200: nil, 201: nil, 204: nil},
		NumTries:          1,
		DelayBetweenTries: 500,
		ResponseLogLimit:  2000,
		Context:           context.Background(),
	}
}

//Returns the *http.Response from the request made. If the request hasn't been 'done' yet, this is nil
func (this *APIRequest) GetHttpResponse() *http.Response {
	return this.HttpResponse
}

//Async wrapper for Do function
func (this *APIRequest) DoAsync(child common.IResult) {
	defer child.Flush()
	child.MergeWithResult(this.Do())
}

//Do is the launch point for the request. It will 'do' the request according to the data that has been set, see each
//data field for more information
func (this *APIRequest) Do() (result common.IResult) {
	result = common.MakeCommonResult(this.config)
	error_count := 0
	result.Debugf("APIRequest.Do called for Method: %s, URL: %s", this.Method, this.Url)

	for error_count < this.NumTries {
		if error_count != 0 {
			time.Sleep(time.Millisecond * time.Duration(this.DelayBetweenTries))
		}
		var req *http.Request
		var req_err error
		if this.RequestBody != nil {
			requestBytes, json_err := json.Marshal(this.RequestBody)
			if json_err != nil {
				result.Errorf("Error marshalling requestBody. Err: %v", json_err)
				return
			}
			result.Debugf("Request Body to send: %s", string(requestBytes))
			req, req_err = http.NewRequest(this.Method, this.Url, bytes.NewBuffer(requestBytes))
			if req_err != nil {
				result.Errorf("Error creating new request. Method: %s Url: %s Err: %v", this.Method, this.Url, req_err)
				error_count++
				continue
			}
			req.Header.Set("Content-Type", "application/json")
		} else {
			req, req_err = http.NewRequest(this.Method, this.Url, nil)
			if req_err != nil {
				result.Errorf("Error creating new request. Method: %s Url: %s Err: %v", this.Method, this.Url, req_err)
				error_count++
				continue
			}
		}

		this.Context = context.WithValue(this.Context, common_routing.CONTEXT_API_NAME, this.ApiName)
		req = req.WithContext(this.Context)

		//Set the headers
		for key, val := range this.Headers {
			req.Header.Set(key, val)
		}

		//Run the request through the passed in func, if any
		if this.Formatter != nil {
			this.Formatter.FormatRequest(req)
		}

		result.DebugMessagef("Starting req. %s", time.Now().Format(time.RFC3339))
		resp, do_err := this.client.Do(req)
		result.DebugMessagef("Finished req. %s", time.Now().Format(time.RFC3339))
		if do_err != nil {
			result.Errorf("Error doing request. req: %s Err: %v", req, do_err)
			error_count++
			continue
		}

		this.HttpResponse = resp

		//Verify we got the expected response code.
		exp_response, ok := this.ValidResponses[resp.StatusCode]
		if !ok {
			err_resp := common.ErrResp{}
			body_bytes, _ := ioutil.ReadAll(resp.Body)
			_ = json.Unmarshal(body_bytes, &err_resp)
			if err_resp.Error == "" { //If the body did not contain an 'error' field. Take the whole resp
				err_resp.Error = string(body_bytes)
			}
			if this.ApiName != "" {
				result.SetResponseMessage(this.ApiName + "_ERROR:" + err_resp.Error)
			} else {
				result.SetResponseMessage("DEPENDENCY_API_ERROR: " + err_resp.Error)
			}
			result.Infof("Bad response code returned for url: %s, valid http responses: %v, "+
				"http response code returned: %d Response Body: %s",
				this.Url,
				this.ValidResponses,
				resp.StatusCode,
				string(body_bytes),
			)
			result.SetStatusCode(resp.StatusCode)
			resp.Body.Close()
			if resp.StatusCode == 401 || resp.StatusCode == 403 {
				return
			}
			error_count++
			continue
		} else {
			result.Debugf("Good Status code returned. StatusCode: %d", this.HttpResponse.StatusCode)
		}

		if exp_response != nil {
			switch asserted_data := exp_response.(type) {
			case *[]byte:
				tmp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					result.Errorf("Bad response.Body returned for url: %s response.Body: %v, err: %s", this.Url, string(tmp), err)
					error_count++
					resp.Body.Close()
					continue
				}
				*asserted_data = tmp
			default:
				//If we're expecting a struct, make sure it's valid by unmarshaling the resp.Body into the ExpectedResponseBody
				body, _ := ioutil.ReadAll(resp.Body)
				json_err := json.Unmarshal(body, exp_response)
				if json_err != nil {
					result.Errorf("Bad response.Body returned for url: %s response.Body: %v, err: %s", this.Url, string(body), json_err)
					error_count++
					resp.Body.Close()
					continue
				}
				if this.config.SafeGetConfigVar("LOGGING_LEVEL") == "DEBUG" ||
					this.config.SafeGetConfigVar("LOGGING_LEVEL") == "DEV" {
					result.Debugf("Response.Body returned: %v", string(body))
				} else if len(string(body)) > this.ResponseLogLimit {
					b := []rune(string(body))
					result.Infof("Large Response.Body returned, showing the first %d chars. (turn debug logs on for more): %v", this.ResponseLogLimit, string(b[:this.ResponseLogLimit]))
				} else {
					result.Infof("Response.Body returned: %v", string(body))
				}
			}

			//If we're supposed to validate the payload, check it
			if v, ok := exp_response.(IPayload); ok {
				if err := v.Valid(); err != nil {
					result.Errorf("ExpectedResponseBody.(IPayload) returned invalid payload with error: %s", err.Error())
					error_count++
					continue
				}
			}
		}

		//No errors encountered
		resp.Body.Close()
		break
	}

	if error_count >= this.NumTries {
		result.Errorf("Unable to get a good response for url: %s", this.Url)
		return
	}

	result.Succeed()
	return
}
