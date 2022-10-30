// Code generated by MockGen. DO NOT EDIT.
// Source: ../interface/remote_use_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"

	entity "github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockExtConnection is a mock of ExtConnection interface.
type MockExtConnection struct {
	ctrl     *gomock.Controller
	recorder *MockExtConnectionMockRecorder
}

// MockExtConnectionMockRecorder is the mock recorder for MockExtConnection.
type MockExtConnectionMockRecorder struct {
	mock *MockExtConnection
}

// NewMockExtConnection creates a new mock instance.
func NewMockExtConnection(ctrl *gomock.Controller) *MockExtConnection {
	mock := &MockExtConnection{ctrl: ctrl}
	mock.recorder = &MockExtConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtConnection) EXPECT() *MockExtConnectionMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockExtConnection) Login(ctx context.Context, login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockExtConnectionMockRecorder) Login(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockExtConnection)(nil).Login), ctx, login, password)
}

// Register mocks base method.
func (m *MockExtConnection) Register(ctx context.Context, login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockExtConnectionMockRecorder) Register(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockExtConnection)(nil).Register), ctx, login, password)
}

// RestoreInfo mocks base method.
func (m *MockExtConnection) RestoreInfo(ctx context.Context, session string, w io.Writer, fileInfo *entity.FileInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreInfo", ctx, session, w, fileInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreInfo indicates an expected call of RestoreInfo.
func (mr *MockExtConnectionMockRecorder) RestoreInfo(ctx, session, w, fileInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreInfo", reflect.TypeOf((*MockExtConnection)(nil).RestoreInfo), ctx, session, w, fileInfo)
}

// StoreInfo mocks base method.
func (m *MockExtConnection) StoreInfo(ctx context.Context, session string, reader io.Reader, info *entity.FileInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreInfo", ctx, session, reader, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreInfo indicates an expected call of StoreInfo.
func (mr *MockExtConnectionMockRecorder) StoreInfo(ctx, session, reader, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreInfo", reflect.TypeOf((*MockExtConnection)(nil).StoreInfo), ctx, session, reader, info)
}

// MockSession is a mock of Session interface.
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession.
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance.
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSession) Get() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSessionMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSession)(nil).Get))
}

// Save mocks base method.
func (m *MockSession) Save(sessionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", sessionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockSessionMockRecorder) Save(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockSession)(nil).Save), sessionID)
}
