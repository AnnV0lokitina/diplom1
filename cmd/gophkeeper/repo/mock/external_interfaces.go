// Code generated by MockGen. DO NOT EDIT.
// Source: ../interface.go

// Package mock is a generated GoMock package.
package mock

import (
	io "io"
	os "os"
	reflect "reflect"

	entity "github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockFileStorageReader is a mock of FileStorageReader interface.
type MockFileStorageReader struct {
	ctrl     *gomock.Controller
	recorder *MockFileStorageReaderMockRecorder
}

// MockFileStorageReaderMockRecorder is the mock recorder for MockFileStorageReader.
type MockFileStorageReaderMockRecorder struct {
	mock *MockFileStorageReader
}

// NewMockFileStorageReader creates a new mock instance.
func NewMockFileStorageReader(ctrl *gomock.Controller) *MockFileStorageReader {
	mock := &MockFileStorageReader{ctrl: ctrl}
	mock.recorder = &MockFileStorageReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStorageReader) EXPECT() *MockFileStorageReaderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockFileStorageReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFileStorageReaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFileStorageReader)(nil).Close))
}

// Empty mocks base method.
func (m *MockFileStorageReader) Empty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Empty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Empty indicates an expected call of Empty.
func (mr *MockFileStorageReaderMockRecorder) Empty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Empty", reflect.TypeOf((*MockFileStorageReader)(nil).Empty))
}

// ReadRecord mocks base method.
func (m *MockFileStorageReader) ReadRecord() (*entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRecord")
	ret0, _ := ret[0].(*entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRecord indicates an expected call of ReadRecord.
func (mr *MockFileStorageReaderMockRecorder) ReadRecord() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRecord", reflect.TypeOf((*MockFileStorageReader)(nil).ReadRecord))
}

// MockFileStorageWriter is a mock of FileStorageWriter interface.
type MockFileStorageWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFileStorageWriterMockRecorder
}

// MockFileStorageWriterMockRecorder is the mock recorder for MockFileStorageWriter.
type MockFileStorageWriterMockRecorder struct {
	mock *MockFileStorageWriter
}

// NewMockFileStorageWriter creates a new mock instance.
func NewMockFileStorageWriter(ctrl *gomock.Controller) *MockFileStorageWriter {
	mock := &MockFileStorageWriter{ctrl: ctrl}
	mock.recorder = &MockFileStorageWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStorageWriter) EXPECT() *MockFileStorageWriterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockFileStorageWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFileStorageWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFileStorageWriter)(nil).Close))
}

// WriteRecord mocks base method.
func (m *MockFileStorageWriter) WriteRecord(record *entity.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRecord", record)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRecord indicates an expected call of WriteRecord.
func (mr *MockFileStorageWriterMockRecorder) WriteRecord(record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRecord", reflect.TypeOf((*MockFileStorageWriter)(nil).WriteRecord), record)
}

// MockFileStorageEnclosure is a mock of FileStorageEnclosure interface.
type MockFileStorageEnclosure struct {
	ctrl     *gomock.Controller
	recorder *MockFileStorageEnclosureMockRecorder
}

// MockFileStorageEnclosureMockRecorder is the mock recorder for MockFileStorageEnclosure.
type MockFileStorageEnclosureMockRecorder struct {
	mock *MockFileStorageEnclosure
}

// NewMockFileStorageEnclosure creates a new mock instance.
func NewMockFileStorageEnclosure(ctrl *gomock.Controller) *MockFileStorageEnclosure {
	mock := &MockFileStorageEnclosure{ctrl: ctrl}
	mock.recorder = &MockFileStorageEnclosureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStorageEnclosure) EXPECT() *MockFileStorageEnclosureMockRecorder {
	return m.recorder
}

// Open mocks base method.
func (m *MockFileStorageEnclosure) Open(fileName string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", fileName)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockFileStorageEnclosureMockRecorder) Open(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockFileStorageEnclosure)(nil).Open), fileName)
}

// Remove mocks base method.
func (m *MockFileStorageEnclosure) Remove(fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockFileStorageEnclosureMockRecorder) Remove(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockFileStorageEnclosure)(nil).Remove), fileName)
}

// Save mocks base method.
func (m *MockFileStorageEnclosure) Save(fileName string, reader io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", fileName, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockFileStorageEnclosureMockRecorder) Save(fileName, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockFileStorageEnclosure)(nil).Save), fileName, reader)
}

// MockArchive is a mock of Archive interface.
type MockArchive struct {
	ctrl     *gomock.Controller
	recorder *MockArchiveMockRecorder
}

// MockArchiveMockRecorder is the mock recorder for MockArchive.
type MockArchiveMockRecorder struct {
	mock *MockArchive
}

// NewMockArchive creates a new mock instance.
func NewMockArchive(ctrl *gomock.Controller) *MockArchive {
	mock := &MockArchive{ctrl: ctrl}
	mock.recorder = &MockArchiveMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArchive) EXPECT() *MockArchiveMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockArchive) GetInfo() (os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockArchiveMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockArchive)(nil).GetInfo))
}

// Pack mocks base method.
func (m *MockArchive) Pack() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pack")
	ret0, _ := ret[0].(error)
	return ret0
}

// Pack indicates an expected call of Pack.
func (mr *MockArchiveMockRecorder) Pack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pack", reflect.TypeOf((*MockArchive)(nil).Pack))
}

// ReadByChunks mocks base method.
func (m *MockArchive) ReadByChunks(w io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByChunks", w)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadByChunks indicates an expected call of ReadByChunks.
func (mr *MockArchiveMockRecorder) ReadByChunks(w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByChunks", reflect.TypeOf((*MockArchive)(nil).ReadByChunks), w)
}

// Unpack mocks base method.
func (m *MockArchive) Unpack() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unpack")
	ret0, _ := ret[0].(error)
	return ret0
}

// Unpack indicates an expected call of Unpack.
func (mr *MockArchiveMockRecorder) Unpack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unpack", reflect.TypeOf((*MockArchive)(nil).Unpack))
}

// WriteByChunks mocks base method.
func (m *MockArchive) WriteByChunks(r io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteByChunks", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteByChunks indicates an expected call of WriteByChunks.
func (mr *MockArchiveMockRecorder) WriteByChunks(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteByChunks", reflect.TypeOf((*MockArchive)(nil).WriteByChunks), r)
}
