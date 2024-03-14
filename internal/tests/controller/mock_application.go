// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/controller/application.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockApplicationController is a mock of ApplicationController interface.
type MockApplicationController struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationControllerMockRecorder
}

// MockApplicationControllerMockRecorder is the mock recorder for MockApplicationController.
type MockApplicationControllerMockRecorder struct {
	mock *MockApplicationController
}

// NewMockApplicationController creates a new mock instance.
func NewMockApplicationController(ctrl *gomock.Controller) *MockApplicationController {
	mock := &MockApplicationController{ctrl: ctrl}
	mock.recorder = &MockApplicationControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationController) EXPECT() *MockApplicationControllerMockRecorder {
	return m.recorder
}

// Handler mocks base method.
func (m *MockApplicationController) Handler(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handler", w, r)
}

// Handler indicates an expected call of Handler.
func (mr *MockApplicationControllerMockRecorder) Handler(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockApplicationController)(nil).Handler), w, r)
}