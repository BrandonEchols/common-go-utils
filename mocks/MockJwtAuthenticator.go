// Code generated by MockGen. DO NOT EDIT.
// Source: routing/JwtAuthenticator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	jwt_go "github.com/dgrijalva/jwt-go"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIJwtAuthenticator is a mock of IJwtAuthenticator interface
type MockIJwtAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockIJwtAuthenticatorMockRecorder
}

// MockIJwtAuthenticatorMockRecorder is the mock recorder for MockIJwtAuthenticator
type MockIJwtAuthenticatorMockRecorder struct {
	mock *MockIJwtAuthenticator
}

// NewMockIJwtAuthenticator creates a new mock instance
func NewMockIJwtAuthenticator(ctrl *gomock.Controller) *MockIJwtAuthenticator {
	mock := &MockIJwtAuthenticator{ctrl: ctrl}
	mock.recorder = &MockIJwtAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIJwtAuthenticator) EXPECT() *MockIJwtAuthenticatorMockRecorder {
	return m.recorder
}

// MakeJWT mocks base method
func (m *MockIJwtAuthenticator) MakeJWT(claims jwt_go.Claims) (string, error) {
	ret := m.ctrl.Call(m, "MakeJWT", claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeJWT indicates an expected call of MakeJWT
func (mr *MockIJwtAuthenticatorMockRecorder) MakeJWT(claims interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeJWT", reflect.TypeOf((*MockIJwtAuthenticator)(nil).MakeJWT), claims)
}

// MakeJWTWithIssuer mocks base method
func (m *MockIJwtAuthenticator) MakeJWTWithIssuer(claims jwt_go.Claims, issuer string) (string, error) {
	ret := m.ctrl.Call(m, "MakeJWTWithIssuer", claims, issuer)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeJWTWithIssuer indicates an expected call of MakeJWTWithIssuer
func (mr *MockIJwtAuthenticatorMockRecorder) MakeJWTWithIssuer(claims, issuer interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeJWTWithIssuer", reflect.TypeOf((*MockIJwtAuthenticator)(nil).MakeJWTWithIssuer), claims, issuer)
}

// MakeJWTWithHeaders mocks base method
func (m *MockIJwtAuthenticator) MakeJWTWithHeaders(claims jwt_go.Claims, headers map[string]string) (string, error) {
	ret := m.ctrl.Call(m, "MakeJWTWithHeaders", claims, headers)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeJWTWithHeaders indicates an expected call of MakeJWTWithHeaders
func (mr *MockIJwtAuthenticatorMockRecorder) MakeJWTWithHeaders(claims, headers interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeJWTWithHeaders", reflect.TypeOf((*MockIJwtAuthenticator)(nil).MakeJWTWithHeaders), claims, headers)
}

// MakeUnsignedJWTWithIssuer mocks base method
func (m *MockIJwtAuthenticator) MakeUnsignedJWTWithIssuer(claims jwt_go.Claims, issuer string) (string, error) {
	ret := m.ctrl.Call(m, "MakeUnsignedJWTWithIssuer", claims, issuer)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeUnsignedJWTWithIssuer indicates an expected call of MakeUnsignedJWTWithIssuer
func (mr *MockIJwtAuthenticatorMockRecorder) MakeUnsignedJWTWithIssuer(claims, issuer interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeUnsignedJWTWithIssuer", reflect.TypeOf((*MockIJwtAuthenticator)(nil).MakeUnsignedJWTWithIssuer), claims, issuer)
}

// MakeUnsignedJWTWithIssuerBase64Encoded mocks base method
func (m *MockIJwtAuthenticator) MakeUnsignedJWTWithIssuerBase64Encoded(claims jwt_go.Claims, issuer string) (string, error) {
	ret := m.ctrl.Call(m, "MakeUnsignedJWTWithIssuerBase64Encoded", claims, issuer)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeUnsignedJWTWithIssuerBase64Encoded indicates an expected call of MakeUnsignedJWTWithIssuerBase64Encoded
func (mr *MockIJwtAuthenticatorMockRecorder) MakeUnsignedJWTWithIssuerBase64Encoded(claims, issuer interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeUnsignedJWTWithIssuerBase64Encoded", reflect.TypeOf((*MockIJwtAuthenticator)(nil).MakeUnsignedJWTWithIssuerBase64Encoded), claims, issuer)
}

// DecodeJwt mocks base method
func (m *MockIJwtAuthenticator) DecodeJwt(fullTokenString string, claims jwt_go.Claims) error {
	ret := m.ctrl.Call(m, "DecodeJwt", fullTokenString, claims)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecodeJwt indicates an expected call of DecodeJwt
func (mr *MockIJwtAuthenticatorMockRecorder) DecodeJwt(fullTokenString, claims interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeJwt", reflect.TypeOf((*MockIJwtAuthenticator)(nil).DecodeJwt), fullTokenString, claims)
}

// DecodeJwtWithKeyFunc mocks base method
func (m *MockIJwtAuthenticator) DecodeJwtWithKeyFunc(fullTokenString string, claims jwt_go.Claims, keyFunc jwt_go.Keyfunc) error {
	ret := m.ctrl.Call(m, "DecodeJwtWithKeyFunc", fullTokenString, claims, keyFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecodeJwtWithKeyFunc indicates an expected call of DecodeJwtWithKeyFunc
func (mr *MockIJwtAuthenticatorMockRecorder) DecodeJwtWithKeyFunc(fullTokenString, claims, keyFunc interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeJwtWithKeyFunc", reflect.TypeOf((*MockIJwtAuthenticator)(nil).DecodeJwtWithKeyFunc), fullTokenString, claims, keyFunc)
}