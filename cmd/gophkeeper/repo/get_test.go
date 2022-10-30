package repo

import (
	"bytes"
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func createTestRecord() *entity.Record {
	return &entity.Record{
		CredentialsList: []entity.Credentials{
			{
				Login:    "login",
				Password: "passwors",
				Meta:     "meta",
			},
		},
		TextFileList: []entity.File{
			{
				Name: "name.txt",
				Meta: "meta",
			},
		},
		BinaryFileList: []entity.File{
			{
				Name: "name.exe",
				Meta: "meta",
			},
		},
		BankCardList: []entity.BankCard{
			{
				Number:     "number",
				ExpDate:    "exp_date",
				Cardholder: "cardholder",
				Code:       "code",
				Meta:       "meta",
			},
		},
	}
}

func TestGet(t *testing.T) {
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
	credList := repo.GetCredentialsList()
	assert.Equal(t, 1, len(credList))
	assert.Equal(t, "login", credList[0].Login)

	cred := repo.GetCredentialsByLogin("incorrect login")
	assert.Nil(t, cred)
	cred = repo.GetCredentialsByLogin("login")
	assert.NotNil(t, cred)
	assert.Equal(t, "meta", cred.Meta)

	// test bank cards
	bankCardList := repo.GetBankCardList()
	assert.Equal(t, 1, len(bankCardList))
	assert.Equal(t, "number", bankCardList[0].Number)

	bankCard := repo.GetBankCardByNumber("incorrect number")
	assert.Nil(t, bankCard)
	bankCard = repo.GetBankCardByNumber("number")
	assert.NotNil(t, bankCard)
	assert.Equal(t, "meta", bankCard.Meta)

	// test text files
	textFileList := repo.GetTextFileList()
	assert.Equal(t, 1, len(textFileList))
	assert.Equal(t, "name.txt", textFileList[0].Name)

	_, _, err := repo.GetTextFileByName("incorrect text file name")
	assert.Equal(t, errorNotFound, err)

	enclosure.EXPECT().Open(gomock.Any()).DoAndReturn(func(_ string) (io.Reader, error) {
		return nil, errors.New("error")
	}).Times(1)
	enclosure.EXPECT().Open(gomock.Any()).DoAndReturn(func(_ string) (io.Reader, error) {
		return &bytes.Buffer{}, nil
	}).Times(1)

	_, _, err = repo.GetTextFileByName("name.txt")
	assert.Error(t, err)
	textFile, _, err := repo.GetTextFileByName("name.txt")
	assert.Nil(t, err)
	assert.NotNil(t, textFile)
	assert.Equal(t, "meta", textFile.Meta)

	// test binary files
	binFileList := repo.GetBinaryFileList()
	assert.Equal(t, 1, len(binFileList))
	assert.Equal(t, "name.exe", binFileList[0].Name)

	_, _, err = repo.GetBinaryFileByName("incorrect text file name")
	assert.Equal(t, errorNotFound, err)

	enclosure.EXPECT().Open(gomock.Any()).DoAndReturn(func(_ string) (io.Reader, error) {
		return nil, errors.New("error")
	}).Times(1)
	enclosure.EXPECT().Open(gomock.Any()).DoAndReturn(func(_ string) (io.Reader, error) {
		return &bytes.Buffer{}, nil
	}).Times(1)

	_, _, err = repo.GetBinaryFileByName("name.exe")
	assert.Error(t, err)
	binFile, _, err := repo.GetBinaryFileByName("name.exe")
	assert.Nil(t, err)
	assert.NotNil(t, binFile)
	assert.Equal(t, "meta", binFile.Meta)
}
