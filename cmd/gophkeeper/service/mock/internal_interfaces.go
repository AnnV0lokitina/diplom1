// Code generated by MockGen. DO NOT EDIT.
// Source: ../interface/remote_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRemote is a mock of Remote interface.
type MockRemote struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteMockRecorder
}

// MockRemoteMockRecorder is the mock recorder for MockRemote.
type MockRemoteMockRecorder struct {
	mock *MockRemote
}

// NewMockRemote creates a new mock instance.
func NewMockRemote(ctrl *gomock.Controller) *MockRemote {
	mock := &MockRemote{ctrl: ctrl}
	mock.recorder = &MockRemoteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemote) EXPECT() *MockRemoteMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockRemote) Login(ctx context.Context, login, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, login, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockRemoteMockRecorder) Login(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockRemote)(nil).Login), ctx, login, password)
}

// ReceiveInfo mocks base method.
func (m *MockRemote) ReceiveInfo(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveInfo", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReceiveInfo indicates an expected call of ReceiveInfo.
func (mr *MockRemoteMockRecorder) ReceiveInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveInfo", reflect.TypeOf((*MockRemote)(nil).ReceiveInfo), ctx)
}

// Register mocks base method.
func (m *MockRemote) Register(ctx context.Context, login, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, login, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockRemoteMockRecorder) Register(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRemote)(nil).Register), ctx, login, password)
}

// SendInfo mocks base method.
func (m *MockRemote) SendInfo(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendInfo", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendInfo indicates an expected call of SendInfo.
func (mr *MockRemoteMockRecorder) SendInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendInfo", reflect.TypeOf((*MockRemote)(nil).SendInfo), ctx)
}