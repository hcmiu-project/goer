// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIGoerClient is a mock of IGoerClient interface.
type MockIGoerClient struct {
	ctrl     *gomock.Controller
	recorder *MockIGoerClientMockRecorder
}

// MockIGoerClientMockRecorder is the mock recorder for MockIGoerClient.
type MockIGoerClientMockRecorder struct {
	mock *MockIGoerClient
}

// NewMockIGoerClient creates a new mock instance.
func NewMockIGoerClient(ctrl *gomock.Controller) *MockIGoerClient {
	mock := &MockIGoerClient{ctrl: ctrl}
	mock.recorder = &MockIGoerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGoerClient) EXPECT() *MockIGoerClientMockRecorder {
	return m.recorder
}

// DeleteSessionId mocks base method.
func (m *MockIGoerClient) DeleteSessionId(domain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionId", domain)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSessionId indicates an expected call of DeleteSessionId.
func (mr *MockIGoerClientMockRecorder) DeleteSessionId(domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionId", reflect.TypeOf((*MockIGoerClient)(nil).DeleteSessionId), domain)
}

// Do mocks base method.
func (m *MockIGoerClient) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockIGoerClientMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockIGoerClient)(nil).Do), req)
}