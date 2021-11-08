// Code generated by MockGen. DO NOT EDIT.
// Source: api_request_factory/APIRequest.go

package mocks

import (
	common "bitbucket.xant.tech/ci/ci-go-utils/common"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockIAPIRequest is a mock of IAPIRequest interface
type MockIAPIRequest struct {
	ctrl     *gomock.Controller
	recorder *MockIAPIRequestMockRecorder
}

// MockIAPIRequestMockRecorder is the mock recorder for MockIAPIRequest
type MockIAPIRequestMockRecorder struct {
	mock *MockIAPIRequest
}

// NewMockIAPIRequest creates a new mock instance
func NewMockIAPIRequest(ctrl *gomock.Controller) *MockIAPIRequest {
	mock := &MockIAPIRequest{ctrl: ctrl}
	mock.recorder = &MockIAPIRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIAPIRequest) EXPECT() *MockIAPIRequestMockRecorder {
	return _m.recorder
}

// DoAsync mocks base method
func (_m *MockIAPIRequest) DoAsync(_param0 common.IResult) {
	_m.ctrl.Call(_m, "DoAsync", _param0)
}

// DoAsync indicates an expected call of DoAsync
func (_mr *MockIAPIRequestMockRecorder) DoAsync(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoAsync", reflect.TypeOf((*MockIAPIRequest)(nil).DoAsync), arg0)
}

// Do mocks base method
func (_m *MockIAPIRequest) Do() common.IResult {
	ret := _m.ctrl.Call(_m, "Do")
	ret0, _ := ret[0].(common.IResult)
	return ret0
}

// Do indicates an expected call of Do
func (_mr *MockIAPIRequestMockRecorder) Do() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Do", reflect.TypeOf((*MockIAPIRequest)(nil).Do))
}

// GetHttpResponse mocks base method
func (_m *MockIAPIRequest) GetHttpResponse() *http.Response {
	ret := _m.ctrl.Call(_m, "GetHttpResponse")
	ret0, _ := ret[0].(*http.Response)
	return ret0
}

// GetHttpResponse indicates an expected call of GetHttpResponse
func (_mr *MockIAPIRequestMockRecorder) GetHttpResponse() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetHttpResponse", reflect.TypeOf((*MockIAPIRequest)(nil).GetHttpResponse))
}

// MockIPayload is a mock of IPayload interface
type MockIPayload struct {
	ctrl     *gomock.Controller
	recorder *MockIPayloadMockRecorder
}

// MockIPayloadMockRecorder is the mock recorder for MockIPayload
type MockIPayloadMockRecorder struct {
	mock *MockIPayload
}

// NewMockIPayload creates a new mock instance
func NewMockIPayload(ctrl *gomock.Controller) *MockIPayload {
	mock := &MockIPayload{ctrl: ctrl}
	mock.recorder = &MockIPayloadMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIPayload) EXPECT() *MockIPayloadMockRecorder {
	return _m.recorder
}

// Valid mocks base method
func (_m *MockIPayload) Valid() error {
	ret := _m.ctrl.Call(_m, "Valid")
	ret0, _ := ret[0].(error)
	return ret0
}

// Valid indicates an expected call of Valid
func (_mr *MockIPayloadMockRecorder) Valid() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Valid", reflect.TypeOf((*MockIPayload)(nil).Valid))
}