package repo

import (
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := mock.NewMockFileStorageReader(ctrl)
	writer := mock.NewMockFileStorageWriter(ctrl)
	enclosure := mock.NewMockFileStorageEnclosure(ctrl)
	archive := mock.NewMockArchive(ctrl)

	record := &entity.Record{}

	repo := &Repo{
		record:    record,
		writer:    writer,
		enclosure: enclosure,
		archive:   archive,
		reader:    reader,
	}
	// test Credentials
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return errors.New("error")
	}).Times(1)
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return nil
	}).Times(1)

	cr := entity.Credentials{
		Login:    "login",
		Password: "password",
		Meta:     "meta",
	}
	err := repo.AddCredentials(cr)
	assert.Error(t, err)
	err = repo.AddCredentials(cr)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.record.CredentialsList))
	assert.Equal(t, "login", repo.record.CredentialsList[0].Login)
	err = repo.AddCredentials(cr)
	assert.Equal(t, errorDuplicate, err)

	// test Bank Card
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return errors.New("error")
	}).Times(1)
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return nil
	}).Times(1)

	card := entity.BankCard{
		Number:     "number",
		ExpDate:    "exp_date",
		Cardholder: "cardholder",
		Code:       "code",
		Meta:       "meta",
	}
	err = repo.AddBankCard(card)
	assert.Error(t, err)
	err = repo.AddBankCard(card)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.record.BankCardList))
	assert.Equal(t, "number", repo.record.BankCardList[0].Number)
	err = repo.AddBankCard(card)
	assert.Equal(t, errorDuplicate, err)

	// test text file
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return errors.New("error")
	}).Times(1)
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return nil
	}).Times(1)
	enclosure.EXPECT().Save(gomock.Any(), gomock.Any()).DoAndReturn(func(_ string, _ io.Reader) error {
		return errors.New("error")
	}).Times(1)
	enclosure.EXPECT().Save(gomock.Any(), gomock.Any()).DoAndReturn(func(_ string, _ io.Reader) error {
		return nil
	}).Times(2)

	textFile := entity.File{
		Name: "text_file.txt",
		Meta: "meta",
	}
	text := "texttexttexttexttexttexttexttexttexttexttexttexttexttexttexttexttexttexttexttext"
	textReader := entity.NewTextReadCloser(text)
	err = repo.AddTextFile(textFile, textReader)
	assert.Error(t, err)
	err = repo.AddTextFile(textFile, textReader)
	assert.Error(t, err)
	err = repo.AddTextFile(textFile, textReader)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.record.TextFileList))
	assert.Equal(t, "text_file.txt", repo.record.TextFileList[0].Name)
	err = repo.AddTextFile(textFile, textReader)
	assert.Equal(t, errorDuplicate, err)

	// test binary file
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return errors.New("error")
	}).Times(1)
	writer.EXPECT().WriteRecord(gomock.Any()).DoAndReturn(func(_ *entity.Record) error {
		return nil
	}).Times(1)
	enclosure.EXPECT().Save(gomock.Any(), gomock.Any()).DoAndReturn(func(_ string, _ io.Reader) error {
		return errors.New("error")
	}).Times(1)
	enclosure.EXPECT().Save(gomock.Any(), gomock.Any()).DoAndReturn(func(_ string, _ io.Reader) error {
		return nil
	}).Times(2)

	binFile := entity.File{
		Name: "bin_file.exe",
		Meta: "meta",
	}
	//binReader := bytes.NewReader([]byte(text))
	binReader := entity.NewTextReadCloser(text)
	err = repo.AddBinaryFile(binFile, binReader)
	assert.Error(t, err)
	err = repo.AddBinaryFile(binFile, binReader)
	assert.Error(t, err)
	err = repo.AddBinaryFile(binFile, binReader)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.record.BinaryFileList))
	assert.Equal(t, "bin_file.exe", repo.record.BinaryFileList[0].Name)
	err = repo.AddBinaryFile(binFile, binReader)
	assert.Equal(t, errorDuplicate, err)
}
