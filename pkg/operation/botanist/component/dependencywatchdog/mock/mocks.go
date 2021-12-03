// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/operation/botanist/component/dependencywatchdog (interfaces: AccessInterface)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccessInterface is a mock of AccessInterface interface.
type MockAccessInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAccessInterfaceMockRecorder
}

// MockAccessInterfaceMockRecorder is the mock recorder for MockAccessInterface.
type MockAccessInterfaceMockRecorder struct {
	mock *MockAccessInterface
}

// NewMockAccessInterface creates a new mock instance.
func NewMockAccessInterface(ctrl *gomock.Controller) *MockAccessInterface {
	mock := &MockAccessInterface{ctrl: ctrl}
	mock.recorder = &MockAccessInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessInterface) EXPECT() *MockAccessInterfaceMockRecorder {
	return m.recorder
}

// Deploy mocks base method.
func (m *MockAccessInterface) Deploy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockAccessInterfaceMockRecorder) Deploy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockAccessInterface)(nil).Deploy), arg0)
}

// Destroy mocks base method.
func (m *MockAccessInterface) Destroy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockAccessInterfaceMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockAccessInterface)(nil).Destroy), arg0)
}

// SetCACertificate mocks base method.
func (m *MockAccessInterface) SetCACertificate(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCACertificate", arg0)
}

// SetCACertificate indicates an expected call of SetCACertificate.
func (mr *MockAccessInterfaceMockRecorder) SetCACertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCACertificate", reflect.TypeOf((*MockAccessInterface)(nil).SetCACertificate), arg0)
}
