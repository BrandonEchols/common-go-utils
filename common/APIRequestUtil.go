/**
 * Copyright 2018 InsideSales.com Inc.
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is the property of InsideSales.com, Inc. and its suppliers, if
 * any. The intellectual and technical concepts contained herein are proprietary and are protected by
 * trade secret or copyright law, and may be covered by U.S. and foreign patents and patents pending.
 * Dissemination of this information or reproduction of this material is strictly forbidden without prior
 * written permission from InsideSales.com Inc.
 *
 * Requests for permission should be addressed to the Legal Department, InsideSales.com,
 * 1712 South East Bay Blvd. Provo, UT 84606.
 *
 * The software and any accompanying documentation are provided "as is" with no warranty.
 * InsideSales.com, Inc. shall not be liable for direct, indirect, special, incidental, consequential, or other
 * damages, under any theory of liability.
 */
package common

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type IAPIRequestUtil interface {
	DoRequest(
		req_formatter IRequestFormatter,
		url string,
		method string,
		requestBody interface{},
	) (*http.Response, IResult)

	RequestWithRetries(
		req_formatter IRequestFormatter,
		url string,
		method string,
		requestBody interface{},
		valid_response_codes []int,
		expected_response_struct interface{},
	) (*http.Response, IResult)

	Request(
		req_formatter IRequestFormatter,
		url string,
		method string,
		requestBody interface{},
		valid_response_codes []int,
		expected_response_struct interface{},
	) (response *http.Response, result IResult)
}

//Implements IAPIRequestUtil
type APIRequestUtils struct {
	client  *http.Client
	configs IConfigGetter
}

func GetAPIRequestUtils(client *http.Client, configs IConfigGetter) IAPIRequestUtil {
	return &APIRequestUtils{
		client:  client,
		configs: configs,
	}
}

//A package-local struct to use to get the error message (if any) from the API response
type ErrResp struct {
	Error string `json:"error"`
}

/*
	DoRequest Makes a request using the parameters passed in.
	@params
		req_formatter IRequestFormatter A function to run the request through prior to sending
		url string The url to make the request to
		method string The http method to use for the request
		requestBody interface The struct to be marshaled and put in the body of the request || nil
	@returns
		*http.Response The response
		IResult The result including logged messages
*/
func (this APIRequestUtils) DoRequest(
	req_formatter IRequestFormatter,
	url string,
	method string,
	requestBody interface{},
) (resp *http.Response, result IResult) {
	result = MakeCommsResult(this.configs)
	result.Infof("DoRequest called with url %s, method %s, requestBody %v", url, method, requestBody)

	var req *http.Request
	var req_err error
	if requestBody != nil {
		requestBytes, json_err := json.Marshal(requestBody)
		if json_err != nil {
			result.Errorf("Error marshalling requestBody in doRequest. Err: %v", json_err)
			return
		}
		req, req_err = http.NewRequest(method, url, bytes.NewBuffer(requestBytes))
		if req_err != nil {
			result.Errorf("Error creating new request. Method: %s Url: %s Err: %v", method, url, req_err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, req_err = http.NewRequest(method, url, nil)
		if req_err != nil {
			result.Errorf("Error creating new request. Method: %s Url: %s Err: %v", method, url, req_err)
			return
		}
	}

	//Run the request through the passed in func
	if req_formatter != nil {
		req_formatter.FormatRequest(req)
	}

	result.DebugMessagef("Starting req. %s", time.Now().Format(time.RFC3339))
	resp, do_err := this.client.Do(req)
	result.DebugMessagef("Finished req. %s", time.Now().Format(time.RFC3339))
	if do_err != nil {
		result.Errorf("Error doing request in doRequest. req: %s Err: %v", req, do_err)
		return
	}

	result.Succeed()
	return
}

/*
	This is a wrapper around the DoRequest function that takes the expected payload and tries up to three times
	to get a valid response with valid data
*/
func (this APIRequestUtils) RequestWithRetries(
	req_formatter IRequestFormatter,
	url string,
	method string,
	requestBody interface{},
	valid_response_codes []int,
	expected_response_struct interface{},
) (response *http.Response, result IResult) {
	result = MakeCommsResult(this.configs)
	error_count := 0

	result.Debugf("RequestWithRetries called for Method: %s, URL: %s", method, url)

	for error_count < 3 { //Try up to three times
		resp, r := this.DoRequest(req_formatter, url, method, requestBody)
		result.MergeWithResult(r)
		//Verify we didn't error on the request.
		if !r.WasSuccessful() {
			error_count++
			continue
		}

		response = resp

		//Verify we got the expected response code.
		valid_http_response := false
		for _, valid_respose_code := range valid_response_codes {
			if resp.StatusCode == valid_respose_code {
				valid_http_response = true
			}
		}
		if valid_http_response == false {
			err_resp := ErrResp{}
			body_bytes, _ := ioutil.ReadAll(resp.Body)
			_ = json.Unmarshal(body_bytes, &err_resp)
			if err_resp.Error == "" { //If the body did not contain an 'error' field. Take the whole resp
				err_resp.Error = string(body_bytes)
			}
			result.SetResponseMessage(err_resp.Error)
			result.Infof("Bad response code returned for url: %s, valid http responses: %v, "+
				"http response code returned: %d Response Body: %s",
				url,
				valid_response_codes,
				resp.StatusCode,
				string(body_bytes),
			)
			result.SetStatusCode(resp.StatusCode)
			resp.Body.Close()
			error_count++
			continue
		}

		if expected_response_struct != nil {
			switch asserted_data := expected_response_struct.(type) {
			case *[]byte:
				tmp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					result.Errorf("Bad response.Body returned for url: %s response.Body: %v, err: %s", url, string(tmp), err)
					error_count++
					resp.Body.Close()
					continue
				}
				*asserted_data = tmp
			default:
				//If we're expecting a struct, make sure it's valid by attempting to unmarshal into the expected_response_struct
				body, _ := ioutil.ReadAll(resp.Body)
				json_err := json.Unmarshal(body, expected_response_struct)
				if json_err != nil {
					result.Errorf("Bad response.Body returned for url: %s response.Body: %v, err: %s", url, string(body), json_err)
					error_count++
					resp.Body.Close()
					continue
				}
				if this.configs.SafeGetConfigVar("LOGGING_LEVEL") == "DEBUG" ||
					this.configs.SafeGetConfigVar("LOGGING_LEVEL") == "DEV" {
					result.Debugf("Response.Body returned: %v", string(body))
				} else if len(string(body)) > 2000 {
					b := []rune(string(body))
					result.Infof("Large Response.Body returned, showing the first 2000 chars. (turn debug logs on for more): %v", string(b[:2000]))
				} else {
					result.Infof("Response.Body returned: %v", string(body))
				}
			}
		}

		//No errors encountered
		resp.Body.Close()
		break
	}

	if error_count >= 2 {
		result.Errorf("Unable to get a good response for url: %s", url)
		return
	}

	result.Succeed()
	return
}

/*
	As RequestWithRetries but only tries once.
*/
func (this APIRequestUtils) Request(
	req_formatter IRequestFormatter,
	url string,
	method string,
	requestBody interface{},
	valid_response_codes []int,
	expected_response_struct interface{},
) (response *http.Response, result IResult) {
	result = MakeCommsResult(this.configs)
	result.Infof("Request called for Method: %s, URL: %s", method, url)

	resp, r := this.DoRequest(req_formatter, url, method, requestBody)
	result.MergeWithResult(r)
	if !r.WasSuccessful() {
		result.Errorf("Error returned from DoRequest for url: %s", url)
		return
	}

	response = resp

	//Verify we got the expected response code.
	valid_http_response := false
	for _, valid_respose_code := range valid_response_codes {
		if resp.StatusCode == valid_respose_code {
			valid_http_response = true
		}
	}
	if valid_http_response == false {
		err_resp := ErrResp{}
		body_bytes, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(body_bytes, &err_resp)
		if err_resp.Error == "" { //If the body did not contain an 'error' field. Take the whole resp
			err_resp.Error = string(body_bytes)
		}
		result.SetResponseMessage(err_resp.Error)
		result.Errorf("Bad response code returned for url: %s, valid http responses: %v, "+
			"http response code returned: %d Response Body: %s",
			url,
			valid_response_codes,
			resp.StatusCode,
			string(body_bytes),
		)
		result.SetStatusCode(resp.StatusCode)
		resp.Body.Close()
		return
	}

	if expected_response_struct != nil {
		switch asserted_data := expected_response_struct.(type) {
		case *[]byte:
			tmp, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				result.Errorf("Bad response.Body returned for url: %s response.Body: %v, err: %s", url, string(tmp), err)
				resp.Body.Close()
				return
			}
			*asserted_data = tmp
		default:
			//If we're expecting a struct, make sure it's valid by attempting to unmarshal into the expected_response_struct
			json_err := json.NewDecoder(resp.Body).Decode(expected_response_struct)
			if json_err != nil {
				body, _ := ioutil.ReadAll(resp.Body)
				result.Errorf("Bad response.Body returned for url: %s response.Body: %v, err: %s", url, string(body), json_err)
				resp.Body.Close()
				return
			}
		}
	}

	//No errors encountered
	resp.Body.Close()
	result.Succeed()
	return
}

type IRequestFormatter interface {
	FormatRequest(*http.Request)
}

type AuthFormatter struct {
	Auth   string //The string to put in the header
	Header string //The Header name. Defaults to Authorization
}

func (this AuthFormatter) FormatRequest(r *http.Request) {
	if this.Header == "" {
		this.Header = "Authorization"
	}
	r.Header.Set(this.Header, this.Auth)
}

type AuthAndContextFormatter struct {
	Auth    string //The string to put in the header
	Header  string //The Header name. Defaults to Authorization
	Context context.Context
}

func (this AuthAndContextFormatter) FormatRequest(r *http.Request) {
	if this.Header == "" {
		this.Header = "Authorization"
	}
	r.Header.Set(this.Header, this.Auth)
	r.WithContext(this.Context)
}

type BasicAppAuthFormatter struct {
	AppId     string //The appUuid to use in the basic authentication
	AppSecret string //The appSecret to use in the basic authentication
	Header    string //The Header name. Defaults to Authorization
}

func (this BasicAppAuthFormatter) FormatRequest(r *http.Request) {
	if this.Header == "" {
		this.Header = "Authorization"
	}
	auth := this.AppId + ":" + this.AppSecret
	r.Header.Add(this.Header, "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))
}

type HeaderRequestFormatterWrapper struct {
	Data      map[string]string //Map of request headers to values
	Formatter IRequestFormatter //Optional
}

func (this HeaderRequestFormatterWrapper) FormatRequest(r *http.Request) {
	for key, val := range this.Data {
		r.Header.Set(key, val)
	}

	if this.Formatter != nil {
		this.Formatter.FormatRequest(r)
	}
}
