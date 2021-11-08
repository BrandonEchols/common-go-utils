// Code generated by MockGen. DO NOT EDIT.
// Source: api_request_factory/APIRequestFactory.go

package mocks

import (
	context "context"
	api_request_factory "github.com/BrandonEchols/common-go-utils/api_request_factory"
	common "github.com/BrandonEchols/common-go-utils/common"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIAPIRequestFactory is a mock of IAPIRequestFactory interface
type MockIAPIRequestFactory struct {
	ctrl     *gomock.Controller
	recorder *MockIAPIRequestFactoryMockRecorder
}

// MockIAPIRequestFactoryMockRecorder is the mock recorder for MockIAPIRequestFactory
type MockIAPIRequestFactoryMockRecorder struct {
	mock *MockIAPIRequestFactory
}

// NewMockIAPIRequestFactory creates a new mock instance
func NewMockIAPIRequestFactory(ctrl *gomock.Controller) *MockIAPIRequestFactory {
	mock := &MockIAPIRequestFactory{ctrl: ctrl}
	mock.recorder = &MockIAPIRequestFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIAPIRequestFactory) EXPECT() *MockIAPIRequestFactoryMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockIAPIRequestFactory) Get(url string, options ...api_request_factory.Opt) api_request_factory.IAPIRequest {
	_s := []interface{}{url}
	for _, _x := range options {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Get", _s...)
	ret0, _ := ret[0].(api_request_factory.IAPIRequest)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockIAPIRequestFactoryMockRecorder) Get(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockIAPIRequestFactory)(nil).Get), _s...)
}

// Delete mocks base method
func (_m *MockIAPIRequestFactory) Delete(url string, options ...api_request_factory.Opt) api_request_factory.IAPIRequest {
	_s := []interface{}{url}
	for _, _x := range options {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Delete", _s...)
	ret0, _ := ret[0].(api_request_factory.IAPIRequest)
	return ret0
}

// Delete indicates an expected call of Delete
func (_mr *MockIAPIRequestFactoryMockRecorder) Delete(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Delete", reflect.TypeOf((*MockIAPIRequestFactory)(nil).Delete), _s...)
}

// Post mocks base method
func (_m *MockIAPIRequestFactory) Post(url string, options ...api_request_factory.Opt) api_request_factory.IAPIRequest {
	_s := []interface{}{url}
	for _, _x := range options {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Post", _s...)
	ret0, _ := ret[0].(api_request_factory.IAPIRequest)
	return ret0
}

// Post indicates an expected call of Post
func (_mr *MockIAPIRequestFactoryMockRecorder) Post(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Post", reflect.TypeOf((*MockIAPIRequestFactory)(nil).Post), _s...)
}

// Patch mocks base method
func (_m *MockIAPIRequestFactory) Patch(url string, options ...api_request_factory.Opt) api_request_factory.IAPIRequest {
	_s := []interface{}{url}
	for _, _x := range options {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Patch", _s...)
	ret0, _ := ret[0].(api_request_factory.IAPIRequest)
	return ret0
}

// Patch indicates an expected call of Patch
func (_mr *MockIAPIRequestFactoryMockRecorder) Patch(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Patch", reflect.TypeOf((*MockIAPIRequestFactory)(nil).Patch), _s...)
}

// Put mocks base method
func (_m *MockIAPIRequestFactory) Put(url string, options ...api_request_factory.Opt) api_request_factory.IAPIRequest {
	_s := []interface{}{url}
	for _, _x := range options {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Put", _s...)
	ret0, _ := ret[0].(api_request_factory.IAPIRequest)
	return ret0
}

// Put indicates an expected call of Put
func (_mr *MockIAPIRequestFactoryMockRecorder) Put(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Put", reflect.TypeOf((*MockIAPIRequestFactory)(nil).Put), _s...)
}

// Url mocks base method
func (_m *MockIAPIRequestFactory) Url(u string) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "Url", u)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// Url indicates an expected call of Url
func (_mr *MockIAPIRequestFactoryMockRecorder) Url(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Url", reflect.TypeOf((*MockIAPIRequestFactory)(nil).Url), arg0)
}

// Method mocks base method
func (_m *MockIAPIRequestFactory) Method(m string) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "Method", m)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// Method indicates an expected call of Method
func (_mr *MockIAPIRequestFactoryMockRecorder) Method(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Method", reflect.TypeOf((*MockIAPIRequestFactory)(nil).Method), arg0)
}

// ApiName mocks base method
func (_m *MockIAPIRequestFactory) ApiName(n string) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "ApiName", n)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// ApiName indicates an expected call of ApiName
func (_mr *MockIAPIRequestFactoryMockRecorder) ApiName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ApiName", reflect.TypeOf((*MockIAPIRequestFactory)(nil).ApiName), arg0)
}

// Headers mocks base method
func (_m *MockIAPIRequestFactory) Headers(h map[string]string) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "Headers", h)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// Headers indicates an expected call of Headers
func (_mr *MockIAPIRequestFactoryMockRecorder) Headers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Headers", reflect.TypeOf((*MockIAPIRequestFactory)(nil).Headers), arg0)
}

// RequestBody mocks base method
func (_m *MockIAPIRequestFactory) RequestBody(i interface{}) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "RequestBody", i)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// RequestBody indicates an expected call of RequestBody
func (_mr *MockIAPIRequestFactoryMockRecorder) RequestBody(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RequestBody", reflect.TypeOf((*MockIAPIRequestFactory)(nil).RequestBody), arg0)
}

// RequestFormatter mocks base method
func (_m *MockIAPIRequestFactory) RequestFormatter(f common.IRequestFormatter) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "RequestFormatter", f)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// RequestFormatter indicates an expected call of RequestFormatter
func (_mr *MockIAPIRequestFactoryMockRecorder) RequestFormatter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RequestFormatter", reflect.TypeOf((*MockIAPIRequestFactory)(nil).RequestFormatter), arg0)
}

// ValidResponses mocks base method
func (_m *MockIAPIRequestFactory) ValidResponses(b map[int]interface{}) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "ValidResponses", b)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// ValidResponses indicates an expected call of ValidResponses
func (_mr *MockIAPIRequestFactoryMockRecorder) ValidResponses(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ValidResponses", reflect.TypeOf((*MockIAPIRequestFactory)(nil).ValidResponses), arg0)
}

// Retry mocks base method
func (_m *MockIAPIRequestFactory) Retry(n int) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "Retry", n)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// Retry indicates an expected call of Retry
func (_mr *MockIAPIRequestFactoryMockRecorder) Retry(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Retry", reflect.TypeOf((*MockIAPIRequestFactory)(nil).Retry), arg0)
}

// DelayBetweenTries mocks base method
func (_m *MockIAPIRequestFactory) DelayBetweenTries(n int) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "DelayBetweenTries", n)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// DelayBetweenTries indicates an expected call of DelayBetweenTries
func (_mr *MockIAPIRequestFactoryMockRecorder) DelayBetweenTries(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DelayBetweenTries", reflect.TypeOf((*MockIAPIRequestFactory)(nil).DelayBetweenTries), arg0)
}

// ResponseLogLimit mocks base method
func (_m *MockIAPIRequestFactory) ResponseLogLimit(n int) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "ResponseLogLimit", n)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// ResponseLogLimit indicates an expected call of ResponseLogLimit
func (_mr *MockIAPIRequestFactoryMockRecorder) ResponseLogLimit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ResponseLogLimit", reflect.TypeOf((*MockIAPIRequestFactory)(nil).ResponseLogLimit), arg0)
}

// WithContext mocks base method
func (_m *MockIAPIRequestFactory) WithContext(c context.Context) api_request_factory.Opt {
	ret := _m.ctrl.Call(_m, "WithContext", c)
	ret0, _ := ret[0].(api_request_factory.Opt)
	return ret0
}

// WithContext indicates an expected call of WithContext
func (_mr *MockIAPIRequestFactoryMockRecorder) WithContext(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WithContext", reflect.TypeOf((*MockIAPIRequestFactory)(nil).WithContext), arg0)
}
