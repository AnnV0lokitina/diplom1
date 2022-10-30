package service

import (
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestAddCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepo(ctrl)
	remote := mock.NewMockRemote(ctrl)
	ctx := context.Background()
	ctxReceiveErr := context.WithValue(ctx, "error", "receive")
	ctxSendErr := context.WithValue(ctx, "error", "send")

	service := &Service{
		r:    remote,
		repo: repo,
	}

	remote.EXPECT().ReceiveInfo(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		if ctx.Value("error") == "receive" {
			return errors.New("receive error")
		}
		return nil
	}).AnyTimes()
	repo.EXPECT().AddCredentials(gomock.Any()).DoAndReturn(func(cred entity.Credentials) error {
		if cred.Login == "login_error" {
			return errors.New("add error")
		}
		return nil
	}).AnyTimes()
	remote.EXPECT().SendInfo(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		if ctx.Value("error") == "send" {
			return errors.New("send error")
		}
		return nil
	}).AnyTimes()

	err := service.AddCredentials(ctxReceiveErr, "login_error", "password", "meta")
	assert.Error(t, err)
	err = service.AddCredentials(ctx, "login_error", "password", "meta")
	assert.Error(t, err)
	err = service.AddCredentials(ctxSendErr, "login", "password", "meta")
	assert.Error(t, err)
	err = service.AddCredentials(ctx, "login", "password", "meta")
	assert.Nil(t, err)
}

func TestAddText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepo(ctrl)
	remote := mock.NewMockRemote(ctrl)
	ctx := context.Background()
	ctxReceiveErr := context.WithValue(ctx, "error", "receive")
	ctxSendErr := context.WithValue(ctx, "error", "send")

	service := &Service{
		r:    remote,
		repo: repo,
	}

	remote.EXPECT().ReceiveInfo(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		if ctx.Value("error") == "receive" {
			return errors.New("receive error")
		}
		return nil
	}).AnyTimes()
	repo.EXPECT().AddTextFile(gomock.Any(), gomock.Any()).DoAndReturn(func(file entity.File, _ io.Reader) error {
		if file.Name == "add_error" {
			return errors.New("add error")
		}
		return nil
	}).AnyTimes()
	remote.EXPECT().SendInfo(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		if ctx.Value("error") == "send" {
			return errors.New("send error")
		}
		return nil
	}).AnyTimes()

	err := service.AddText(ctxReceiveErr, "text", "add_error", "meta")
	assert.Error(t, err)
	err = service.AddText(ctx, "text", "add_error", "meta")
	assert.Error(t, err)
	err = service.AddText(ctxSendErr, "text", "add", "meta")
	assert.Error(t, err)
	err = service.AddText(ctx, "text", "add", "meta")
	assert.Nil(t, err)
}

func TestAddBinaryDataFromFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepo(ctrl)
	remote := mock.NewMockRemote(ctrl)
	ext := mock.NewMockExternal(ctrl)
	ctx := context.Background()
	ctxReceiveErr := context.WithValue(ctx, "error", "receive")
	ctxSendErr := context.WithValue(ctx, "error", "send")

	service := &Service{
		r:    remote,
		repo: repo,
		ext:  ext,
	}

	remote.EXPECT().ReceiveInfo(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		if ctx.Value("error") == "receive" {
			return errors.New("receive error")
		}
		return nil
	}).AnyTimes()
	ext.EXPECT().Open(gomock.Any()).DoAndReturn(func(filePath string) (string, io.ReadCloser, error) {
		if filePath == "open_external_path" {
			return "", nil, errors.New("open external error")
		}
		r := entity.NewTextReadCloser("i am text")
		return filePath, r, nil
	}).AnyTimes()
	repo.EXPECT().AddBinaryFile(gomock.Any(), gomock.Any()).DoAndReturn(func(file entity.File, _ io.Reader) error {
		if file.Name == "add_error_path" {
			return errors.New("add error")
		}
		return nil
	}).AnyTimes()
	remote.EXPECT().SendInfo(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		if ctx.Value("error") == "send" {
			return errors.New("send error")
		}
		return nil
	}).AnyTimes()

	err := service.AddBinaryDataFromFile(ctxReceiveErr, "open_external_path", "meta")
	assert.Error(t, err)
	err = service.AddBinaryDataFromFile(ctx, "open_external_path", "meta")
	assert.Error(t, err)
	err = service.AddBinaryDataFromFile(ctx, "add_error_path", "meta")
	assert.Error(t, err)
	err = service.AddBinaryDataFromFile(ctxSendErr, "path", "meta")
	assert.Error(t, err)
	err = service.AddBinaryDataFromFile(ctx, "path", "meta")
	assert.Nil(t, err)
}

func TestAddBankCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepo(ctrl)
	remote := mock.NewMockRemote(ctrl)
	ctx := context.Background()
	ctxReceiveErr := context.WithValue(ctx, "error", "receive")
	ctxSendErr := context.WithValue(ctx, "error", "send")

	service := &Service{
		r:    remote,
		repo: repo,
	}

	remote.EXPECT().ReceiveInfo(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		if ctx.Value("error") == "receive" {
			return errors.New("receive error")
		}
		return nil
	}).AnyTimes()
	repo.EXPECT().AddBankCard(gomock.Any()).DoAndReturn(func(card entity.BankCard) error {
		if card.Number == "add_error" {
			return errors.New("add error")
		}
		return nil
	}).AnyTimes()
	remote.EXPECT().SendInfo(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		if ctx.Value("error") == "send" {
			return errors.New("send error")
		}
		return nil
	}).AnyTimes()

	err := service.AddBankCard(ctxReceiveErr, "add_error", "t", "t", "t", "t")
	assert.Error(t, err)
	err = service.AddBankCard(ctx, "add_error", "t", "t", "t", "t")
	assert.Error(t, err)
	err = service.AddBankCard(ctxSendErr, "num", "t", "t", "t", "t")
	assert.Error(t, err)
	err = service.AddBankCard(ctx, "num", "t", "t", "t", "t")
	assert.Nil(t, err)
}
