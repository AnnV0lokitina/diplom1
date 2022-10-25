package service

import (
	"bufio"
	"context"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"io"
)

type Repo interface {
	AddTextFile(file entity.File, reader *bufio.Reader) error
	AddBinaryFile(file entity.File, reader *bufio.Reader) error
	AddCredentials(cred entity.Credentials) error
	AddBankCard(card entity.BankCard) error
	GetTextFileList() []entity.File
	GetTextFileByName(name string) (*entity.File, io.Reader, error)
	GetBinaryFileList() []entity.File
	GetBinaryFileByName(name string) (*entity.File, io.Reader, error)
	GetCredentialsList() []entity.Credentials
	GetCredentialsByLogin(login string) *entity.Credentials
	GetBankCardList() []entity.BankCard
	GetBankCardByNumber(number string) *entity.BankCard
	RemoveTextFileByName(name string) error
	RemoveBinaryFileByName(name string) error
	RemoveCredentialsByLogin(login string) error
	RemoveBankCardByNumber(number string) error
	CreateZIP() error
	UnpackZIP() error
	GetInfo() (*entity.FileInfo, error)
	ReadFileByChunks(w io.Writer) error
	WriteFileByChunks(reader io.Reader) error
}

type ExtConnection interface {
	Register(ctx context.Context, login string, password string) (string, error)
	Login(ctx context.Context, login string, password string) (string, error)
	StoreInfo(ctx context.Context, session string, reader io.Reader, info *entity.FileInfo) error
	RestoreInfo(ctx context.Context, session string, w io.Writer, fileInfo *entity.FileInfo) error
}

type Session interface {
	Save(sessionID string) error
	Get() (string, error)
}
