// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/requester/requester.go

// Package mock_requester is a generated GoMock package.
package mock_requester

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpRequest is a mock of HttpRequest interface.
type MockHttpRequest struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRequestMockRecorder
}

// MockHttpRequestMockRecorder is the mock recorder for MockHttpRequest.
type MockHttpRequestMockRecorder struct {
	mock *MockHttpRequest
}

// NewMockHttpRequest creates a new mock instance.
func NewMockHttpRequest(ctrl *gomock.Controller) *MockHttpRequest {
	mock := &MockHttpRequest{ctrl: ctrl}
	mock.recorder = &MockHttpRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpRequest) EXPECT() *MockHttpRequestMockRecorder {
	return m.recorder
}

// MakeRequest mocks base method.
func (m *MockHttpRequest) MakeRequest(url string) (string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRequest", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MakeRequest indicates an expected call of MakeRequest.
func (mr *MockHttpRequestMockRecorder) MakeRequest(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*MockHttpRequest)(nil).MakeRequest), url)
}

// Normalize mocks base method.
func (m *MockHttpRequest) Normalize(str string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Normalize", str)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Normalize indicates an expected call of Normalize.
func (mr *MockHttpRequestMockRecorder) Normalize(str interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Normalize", reflect.TypeOf((*MockHttpRequest)(nil).Normalize), str)
}