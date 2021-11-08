package api_request_factory

import (
	"context"
	"github.com/BrandonEchols/common-go-utils/common"
)

//Opt's are wrapper functions that can represent anything and everything you can do to customize a typical request
type Opt func(*APIRequest) *APIRequest

func (this apiRequestFactory) Url(u string) Opt {
	return func(a *APIRequest) *APIRequest {
		a.Url = u
		return a
	}
}
func (this apiRequestFactory) Method(m string) Opt {
	return func(a *APIRequest) *APIRequest {
		a.Method = m
		return a
	}
}
func (this apiRequestFactory) ApiName(n string) Opt {
	return func(a *APIRequest) *APIRequest {
		a.ApiName = n
		return a
	}
}
func (this apiRequestFactory) Headers(h map[string]string) Opt {
	return func(a *APIRequest) *APIRequest {
		a.Headers = h
		return a
	}
}
func (this apiRequestFactory) RequestBody(i interface{}) Opt {
	return func(a *APIRequest) *APIRequest {
		a.RequestBody = i
		return a
	}
}
func (this apiRequestFactory) RequestFormatter(f common.IRequestFormatter) Opt {
	return func(a *APIRequest) *APIRequest {
		a.Formatter = f
		return a
	}
}
func (this apiRequestFactory) ValidResponses(b map[int]interface{}) Opt {
	return func(a *APIRequest) *APIRequest {
		a.ValidResponses = b
		return a
	}
}
func (this apiRequestFactory) Retry(n int) Opt {
	return func(a *APIRequest) *APIRequest {
		a.NumTries = n
		return a
	}
}
func (this apiRequestFactory) DelayBetweenTries(n int) Opt {
	return func(a *APIRequest) *APIRequest {
		a.DelayBetweenTries = n
		return a
	}
}
func (this apiRequestFactory) ResponseLogLimit(n int) Opt {
	return func(a *APIRequest) *APIRequest {
		a.ResponseLogLimit = n
		return a
	}
}
func (this apiRequestFactory) WithContext(c context.Context) Opt {
	return func(a *APIRequest) *APIRequest {
		a.Context = c
		return a
	}
}
