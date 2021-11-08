/**
 * Copyright 2018-2019 InsideSales.com Inc.
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
package api_request_factory

import (
	"bitbucket.xant.tech/ci/ci-go-utils/common"
	"context"
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
