package service

import (
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/interface"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepo(ctrl)
	conn := mock.NewMockExtConnection(ctrl)
	s := mock.NewMockSession(ctrl)
	ext := mock.NewMockExternal(ctrl)

	service := NewService(repo, conn, s, ext)
	_, ok := service.repo.(_interface.Repo)
	assert.True(t, ok)
	_, ok = service.r.(_interface.Remote)
	assert.True(t, ok)
}

func TestUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepo(ctrl)
	remote := mock.NewMockRemote(ctrl)
	ctx := context.Background()

	service := &Service{
		r:    remote,
		repo: repo,
	}

	remote.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
	remote.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	remote.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
	remote.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err := service.Register(ctx, "login", "password")
	assert.Error(t, err)
	err = service.Register(ctx, "login", "password")
	assert.Nil(t, err)

	err = service.Login(ctx, "login", "password")
	assert.Error(t, err)
	err = service.Login(ctx, "login", "password")
	assert.Nil(t, err)
}
