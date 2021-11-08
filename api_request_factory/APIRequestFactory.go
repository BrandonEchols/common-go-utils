package api_request_factory

import (
	"context"
	"github.com/BrandonEchols/common-go-utils/common"
	"net/http"
)

type IAPIRequestFactory interface {
	Get(url string, options ...Opt) IAPIRequest
	Delete(url string, options ...Opt) IAPIRequest
	Post(url string, options ...Opt) IAPIRequest
	Patch(url string, options ...Opt) IAPIRequest
	Put(url string, options ...Opt) IAPIRequest

	//Common Options
	Url(u string) Opt
	Method(m string) Opt
	ApiName(n string) Opt
	Headers(h map[string]string) Opt
	RequestBody(i interface{}) Opt
	RequestFormatter(f common.IRequestFormatter) Opt
	ValidResponses(b map[int]interface{}) Opt
	Retry(n int) Opt
	DelayBetweenTries(n int) Opt
	ResponseLogLimit(n int) Opt
	WithContext(c context.Context) Opt
}

const X_REQUEST_SOURCE_HEADER = "x-request-source"

//Implements IAPIRequestFactory
type apiRequestFactory struct {
	client         *http.Client
	config         common.IConfigGetter
	request_source string
}

/*
	@params
		client *http.Client The http client to use for requests
		config common.IConfigGetter The config getter to use for needed configs
		request_source string The string to put in the 'x-request-source' header
*/
func GetAPIRequestFactory(client *http.Client, config common.IConfigGetter, request_source string) IAPIRequestFactory {
	return &apiRequestFactory{
		client:         client,
		config:         config,
		request_source: request_source,
	}
}

/*
	The following functions are almost Identical. The name of the function corresponds to the http METHOD of the request
	It's worth noting that there are defaults that are set for requests. They can be overridden with Opt's passed in.
	See the documentation for getBaseAPIRequest for more info on defaults used
	@Params
		url string This is the full URL to make the request to
		options ...Opt An arbitrary number of 'Opt' (options) can be passed in to alter the request.
			See APIRequestOptions.go for more info.
*/

func (this *apiRequestFactory) Get(url string, options ...Opt) IAPIRequest {
	r := getBaseAPIRequest(this.client, this.config, url, "GET")
	r.Headers[X_REQUEST_SOURCE_HEADER] = this.request_source
	for _, opt := range options {
		r = opt(r)
	}
	return r
}
func (this *apiRequestFactory) Delete(url string, options ...Opt) IAPIRequest {
	r := getBaseAPIRequest(this.client, this.config, url, "DELETE")
	r.Headers[X_REQUEST_SOURCE_HEADER] = this.request_source
	for _, opt := range options {
		r = opt(r)
	}
	return r
}
func (this *apiRequestFactory) Post(url string, options ...Opt) IAPIRequest {
	r := getBaseAPIRequest(this.client, this.config, url, "POST")
	r.Headers[X_REQUEST_SOURCE_HEADER] = this.request_source
	for _, opt := range options {
		r = opt(r)
	}
	return r
}
func (this *apiRequestFactory) Patch(url string, options ...Opt) IAPIRequest {
	r := getBaseAPIRequest(this.client, this.config, url, "PATCH")
	r.Headers[X_REQUEST_SOURCE_HEADER] = this.request_source
	for _, opt := range options {
		r = opt(r)
	}
	return r
}
func (this *apiRequestFactory) Put(url string, options ...Opt) IAPIRequest {
	r := getBaseAPIRequest(this.client, this.config, url, "PUT")
	r.Headers[X_REQUEST_SOURCE_HEADER] = this.request_source
	for _, opt := range options {
		r = opt(r)
	}
	return r
}
