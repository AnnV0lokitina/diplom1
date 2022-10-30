// Code generated by MockGen. DO NOT EDIT.
// Source: ../interface/repo_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	io "io"
	reflect "reflect"

	entity "github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddBankCard mocks base method.
func (m *MockRepo) AddBankCard(card entity.BankCard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBankCard", card)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBankCard indicates an expected call of AddBankCard.
func (mr *MockRepoMockRecorder) AddBankCard(card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBankCard", reflect.TypeOf((*MockRepo)(nil).AddBankCard), card)
}

// AddBinaryFile mocks base method.
func (m *MockRepo) AddBinaryFile(file entity.File, reader io.ReadCloser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBinaryFile", file, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBinaryFile indicates an expected call of AddBinaryFile.
func (mr *MockRepoMockRecorder) AddBinaryFile(file, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinaryFile", reflect.TypeOf((*MockRepo)(nil).AddBinaryFile), file, reader)
}

// AddCredentials mocks base method.
func (m *MockRepo) AddCredentials(cred entity.Credentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCredentials", cred)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCredentials indicates an expected call of AddCredentials.
func (mr *MockRepoMockRecorder) AddCredentials(cred interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCredentials", reflect.TypeOf((*MockRepo)(nil).AddCredentials), cred)
}

// AddTextFile mocks base method.
func (m *MockRepo) AddTextFile(file entity.File, reader io.ReadCloser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTextFile", file, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTextFile indicates an expected call of AddTextFile.
func (mr *MockRepoMockRecorder) AddTextFile(file, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTextFile", reflect.TypeOf((*MockRepo)(nil).AddTextFile), file, reader)
}

// CreateZIP mocks base method.
func (m *MockRepo) CreateZIP() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateZIP")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateZIP indicates an expected call of CreateZIP.
func (mr *MockRepoMockRecorder) CreateZIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateZIP", reflect.TypeOf((*MockRepo)(nil).CreateZIP))
}

// GetBankCardByNumber mocks base method.
func (m *MockRepo) GetBankCardByNumber(number string) *entity.BankCard {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBankCardByNumber", number)
	ret0, _ := ret[0].(*entity.BankCard)
	return ret0
}

// GetBankCardByNumber indicates an expected call of GetBankCardByNumber.
func (mr *MockRepoMockRecorder) GetBankCardByNumber(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBankCardByNumber", reflect.TypeOf((*MockRepo)(nil).GetBankCardByNumber), number)
}

// GetBankCardList mocks base method.
func (m *MockRepo) GetBankCardList() []entity.BankCard {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBankCardList")
	ret0, _ := ret[0].([]entity.BankCard)
	return ret0
}

// GetBankCardList indicates an expected call of GetBankCardList.
func (mr *MockRepoMockRecorder) GetBankCardList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBankCardList", reflect.TypeOf((*MockRepo)(nil).GetBankCardList))
}

// GetBinaryFileByName mocks base method.
func (m *MockRepo) GetBinaryFileByName(name string) (*entity.File, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBinaryFileByName", name)
	ret0, _ := ret[0].(*entity.File)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBinaryFileByName indicates an expected call of GetBinaryFileByName.
func (mr *MockRepoMockRecorder) GetBinaryFileByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinaryFileByName", reflect.TypeOf((*MockRepo)(nil).GetBinaryFileByName), name)
}

// GetBinaryFileList mocks base method.
func (m *MockRepo) GetBinaryFileList() []entity.File {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBinaryFileList")
	ret0, _ := ret[0].([]entity.File)
	return ret0
}

// GetBinaryFileList indicates an expected call of GetBinaryFileList.
func (mr *MockRepoMockRecorder) GetBinaryFileList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinaryFileList", reflect.TypeOf((*MockRepo)(nil).GetBinaryFileList))
}

// GetCredentialsByLogin mocks base method.
func (m *MockRepo) GetCredentialsByLogin(login string) *entity.Credentials {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialsByLogin", login)
	ret0, _ := ret[0].(*entity.Credentials)
	return ret0
}

// GetCredentialsByLogin indicates an expected call of GetCredentialsByLogin.
func (mr *MockRepoMockRecorder) GetCredentialsByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsByLogin", reflect.TypeOf((*MockRepo)(nil).GetCredentialsByLogin), login)
}

// GetCredentialsList mocks base method.
func (m *MockRepo) GetCredentialsList() []entity.Credentials {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialsList")
	ret0, _ := ret[0].([]entity.Credentials)
	return ret0
}

// GetCredentialsList indicates an expected call of GetCredentialsList.
func (mr *MockRepoMockRecorder) GetCredentialsList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsList", reflect.TypeOf((*MockRepo)(nil).GetCredentialsList))
}

// GetInfo mocks base method.
func (m *MockRepo) GetInfo() (*entity.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(*entity.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockRepoMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockRepo)(nil).GetInfo))
}

// GetTextFileByName mocks base method.
func (m *MockRepo) GetTextFileByName(name string) (*entity.File, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTextFileByName", name)
	ret0, _ := ret[0].(*entity.File)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTextFileByName indicates an expected call of GetTextFileByName.
func (mr *MockRepoMockRecorder) GetTextFileByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTextFileByName", reflect.TypeOf((*MockRepo)(nil).GetTextFileByName), name)
}

// GetTextFileList mocks base method.
func (m *MockRepo) GetTextFileList() []entity.File {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTextFileList")
	ret0, _ := ret[0].([]entity.File)
	return ret0
}

// GetTextFileList indicates an expected call of GetTextFileList.
func (mr *MockRepoMockRecorder) GetTextFileList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTextFileList", reflect.TypeOf((*MockRepo)(nil).GetTextFileList))
}

// ReadFileByChunks mocks base method.
func (m *MockRepo) ReadFileByChunks(w io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFileByChunks", w)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadFileByChunks indicates an expected call of ReadFileByChunks.
func (mr *MockRepoMockRecorder) ReadFileByChunks(w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFileByChunks", reflect.TypeOf((*MockRepo)(nil).ReadFileByChunks), w)
}

// RemoveBankCardByNumber mocks base method.
func (m *MockRepo) RemoveBankCardByNumber(number string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBankCardByNumber", number)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveBankCardByNumber indicates an expected call of RemoveBankCardByNumber.
func (mr *MockRepoMockRecorder) RemoveBankCardByNumber(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBankCardByNumber", reflect.TypeOf((*MockRepo)(nil).RemoveBankCardByNumber), number)
}

// RemoveBinaryFileByName mocks base method.
func (m *MockRepo) RemoveBinaryFileByName(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBinaryFileByName", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveBinaryFileByName indicates an expected call of RemoveBinaryFileByName.
func (mr *MockRepoMockRecorder) RemoveBinaryFileByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBinaryFileByName", reflect.TypeOf((*MockRepo)(nil).RemoveBinaryFileByName), name)
}

// RemoveCredentialsByLogin mocks base method.
func (m *MockRepo) RemoveCredentialsByLogin(login string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCredentialsByLogin", login)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCredentialsByLogin indicates an expected call of RemoveCredentialsByLogin.
func (mr *MockRepoMockRecorder) RemoveCredentialsByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCredentialsByLogin", reflect.TypeOf((*MockRepo)(nil).RemoveCredentialsByLogin), login)
}

// RemoveTextFileByName mocks base method.
func (m *MockRepo) RemoveTextFileByName(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTextFileByName", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTextFileByName indicates an expected call of RemoveTextFileByName.
func (mr *MockRepoMockRecorder) RemoveTextFileByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTextFileByName", reflect.TypeOf((*MockRepo)(nil).RemoveTextFileByName), name)
}

// UnpackZIP mocks base method.
func (m *MockRepo) UnpackZIP() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpackZIP")
	ret0, _ := ret[0].(error)
	return ret0
}

// UnpackZIP indicates an expected call of UnpackZIP.
func (mr *MockRepoMockRecorder) UnpackZIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpackZIP", reflect.TypeOf((*MockRepo)(nil).UnpackZIP))
}

// WriteFileByChunks mocks base method.
func (m *MockRepo) WriteFileByChunks(reader io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFileByChunks", reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFileByChunks indicates an expected call of WriteFileByChunks.
func (mr *MockRepoMockRecorder) WriteFileByChunks(reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFileByChunks", reflect.TypeOf((*MockRepo)(nil).WriteFileByChunks), reader)
}