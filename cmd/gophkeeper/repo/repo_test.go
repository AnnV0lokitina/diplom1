package repo

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFileRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := mock.NewMockFileStorageReader(ctrl)
	writer := mock.NewMockFileStorageWriter(ctrl)
	enclosure := mock.NewMockFileStorageEnclosure(ctrl)
	archive := mock.NewMockArchive(ctrl)

	reader.EXPECT().Empty().DoAndReturn(func() bool {
		return true
	}).Times(1)
	writer.EXPECT().Close().DoAndReturn(func() error {
		return nil
	}).Times(1)

	r, err := NewFileRepo(reader, writer, enclosure, archive)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(r.record.TextFileList))
	assert.Equal(t, 0, len(r.record.CredentialsList))
	assert.Equal(t, 0, len(r.record.BinaryFileList))
	assert.Equal(t, 0, len(r.record.BankCardList))
	err = r.Close()
	assert.Nil(t, err)

	assert.Equal(t, archive, r.archive)
	assert.Equal(t, writer, r.writer)
	assert.Equal(t, enclosure, r.enclosure)

	reader.EXPECT().Empty().DoAndReturn(func() bool {
		return false
	}).Times(1)
	reader.EXPECT().ReadRecord().DoAndReturn(func() (*entity.Record, error) {
		return &entity.Record{
			CredentialsList: []entity.Credentials{
				{
					Login:    "login",
					Password: "password",
					Meta:     "meta",
				},
			},
		}, nil
	}).Times(1)
	writer.EXPECT().Close().DoAndReturn(func() error {
		return nil
	}).Times(1)
	reader.EXPECT().Close().DoAndReturn(func() error {
		return nil
	}).Times(1)

	r, err = NewFileRepo(reader, writer, enclosure, archive)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(r.record.TextFileList))
	assert.Equal(t, 1, len(r.record.CredentialsList))
	assert.Equal(t, 0, len(r.record.BinaryFileList))
	assert.Equal(t, 0, len(r.record.BankCardList))
	err = r.Close()
	assert.Nil(t, err)
}
