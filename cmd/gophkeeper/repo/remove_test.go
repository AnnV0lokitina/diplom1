package repo

import (
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := mock.NewMockFileStorageReader(ctrl)
	writer := mock.NewMockFileStorageWriter(ctrl)
	enclosure := mock.NewMockFileStorageEnclosure(ctrl)
	archive := mock.NewMockArchive(ctrl)

	record := createTestRecord()

	repo := &Repo{
		record:    record,
		writer:    writer,
		enclosure: enclosure,
		archive:   archive,
		reader:    reader,
	}

	// test Credentials
	assert.Equal(t, 1, len(repo.record.CredentialsList))
	err := repo.RemoveCredentialsByLogin(repo.record.CredentialsList[0].Login + "1")
	assert.Equal(t, errorNotFound, err)

	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return errors.New("error")
	}).Times(1)
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return nil
	}).Times(1)

	err = repo.RemoveCredentialsByLogin(repo.record.CredentialsList[0].Login)
	assert.Error(t, err)
	assert.Equal(t, 1, len(repo.record.CredentialsList))

	err = repo.RemoveCredentialsByLogin(repo.record.CredentialsList[0].Login)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(repo.record.CredentialsList))

	// test bank card
	assert.Equal(t, 1, len(repo.record.BankCardList))
	err = repo.RemoveBankCardByNumber(repo.record.BankCardList[0].Number + "1")
	assert.Equal(t, errorNotFound, err)

	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return errors.New("error")
	}).Times(1)
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return nil
	}).Times(1)

	err = repo.RemoveBankCardByNumber(repo.record.BankCardList[0].Number)
	assert.Error(t, err)
	assert.Equal(t, 1, len(repo.record.BankCardList))

	err = repo.RemoveBankCardByNumber(repo.record.BankCardList[0].Number)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(repo.record.BankCardList))

	// test text file
	assert.Equal(t, 1, len(repo.record.TextFileList))
	err = repo.RemoveTextFileByName(repo.record.TextFileList[0].Name + "1")
	assert.Equal(t, errorNotFound, err)

	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return errors.New("error")
	}).Times(1)
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return nil
	}).Times(1)

	enclosure.EXPECT().Remove(gomock.Any()).DoAndReturn(func(_ string) error {
		return errors.New("error")
	}).Times(1)
	enclosure.EXPECT().Remove(gomock.Any()).DoAndReturn(func(_ string) error {
		return nil
	}).Times(2)

	err = repo.RemoveTextFileByName(repo.record.TextFileList[0].Name)
	assert.Error(t, err)
	assert.Equal(t, 1, len(repo.record.TextFileList))

	err = repo.RemoveTextFileByName(repo.record.TextFileList[0].Name)
	assert.Error(t, err)
	assert.Equal(t, 1, len(repo.record.TextFileList))

	err = repo.RemoveTextFileByName(repo.record.TextFileList[0].Name)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(repo.record.TextFileList))

	// test binary file
	assert.Equal(t, 1, len(repo.record.BinaryFileList))
	err = repo.RemoveBinaryFileByName(repo.record.BinaryFileList[0].Name + "1")
	assert.Equal(t, errorNotFound, err)

	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return errors.New("error")
	}).Times(1)
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return nil
	}).Times(1)

	enclosure.EXPECT().Remove(gomock.Any()).DoAndReturn(func(_ string) error {
		return errors.New("error")
	}).Times(1)
	enclosure.EXPECT().Remove(gomock.Any()).DoAndReturn(func(_ string) error {
		return nil
	}).Times(2)

	err = repo.RemoveBinaryFileByName(repo.record.BinaryFileList[0].Name)
	assert.Error(t, err)
	assert.Equal(t, 1, len(repo.record.BinaryFileList))

	err = repo.RemoveBinaryFileByName(repo.record.BinaryFileList[0].Name)
	assert.Error(t, err)
	assert.Equal(t, 1, len(repo.record.BinaryFileList))

	err = repo.RemoveBinaryFileByName(repo.record.BinaryFileList[0].Name)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(repo.record.BinaryFileList))
}
