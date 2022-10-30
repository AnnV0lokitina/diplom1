package _interface

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"io"
)

type RepoData interface {
	AddTextFile(file entity.File, reader io.ReadCloser) error
	AddBinaryFile(file entity.File, reader io.ReadCloser) error
	AddCredentials(cred entity.Credentials) error
	AddBankCard(card entity.BankCard) error
	GetTextFileList() []entity.File
	GetTextFileByName(name string) (*entity.File, io.ReadCloser, error)
	GetBinaryFileList() []entity.File
	GetBinaryFileByName(name string) (*entity.File, io.ReadCloser, error)
	GetCredentialsList() []entity.Credentials
	GetCredentialsByLogin(login string) *entity.Credentials
	GetBankCardList() []entity.BankCard
	GetBankCardByNumber(number string) *entity.BankCard
	RemoveTextFileByName(name string) error
	RemoveBinaryFileByName(name string) error
	RemoveCredentialsByLogin(login string) error
	RemoveBankCardByNumber(number string) error
}
