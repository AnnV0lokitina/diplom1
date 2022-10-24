package handler

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
)

// ServiceAdd declare add interface
type ServiceAdd interface {
	AddCredentials(ctx context.Context, login string, password string, meta string) error
	AddTextFromFile(ctx context.Context, path string, meta string) error
	AddBinaryDataFromFile(ctx context.Context, path string, meta string) error
	AddBankCard(
		ctx context.Context,
		number string,
		exp string,
		cardholder string,
		code string,
		meta string,
	) error
}

// ServiceReg declare login/registration interface
type ServiceReg interface {
	Register(ctx context.Context, login string, password string) error
	Login(ctx context.Context, login string, password string) error
}

// ServiceRemove declare remove interface
type ServiceRemove interface {
	RemoveCredentialsByLogin(ctx context.Context, login string) error
	RemoveTextByName(ctx context.Context, name string) error
	RemoveBinaryDataByName(ctx context.Context, name string) error
	RemoveBankCardByNumber(ctx context.Context, number string) error
}

// ServiceShow declare show interface
type ServiceShow interface {
	ShowCredentialsList(ctx context.Context) []entity.Credentials
	ShowTextFilesList(ctx context.Context) []entity.File
	ShowBinaryDataList(ctx context.Context) []entity.File
	ShowBankCardList(ctx context.Context) []entity.BankCard
	GetCredentialsByLogin(ctx context.Context, login string) *entity.Credentials
	GetBankCardByNumber(ctx context.Context, number string) *entity.BankCard
	UploadTextFileByNameIntoPath(ctx context.Context, name string, outFilePath string) (*entity.File, error)
	UploadBinaryFileByNameIntoPath(ctx context.Context, name string, outFilePath string) (*entity.File, error)
}

// Service declare Service interface
type Service interface {
	ServiceAdd
	ServiceReg
	ServiceRemove
	ServiceShow
}
