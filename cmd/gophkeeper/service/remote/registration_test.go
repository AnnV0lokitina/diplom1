package remote

import (
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/remote/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepoZip(ctrl)
	conn := mock.NewMockExtConnection(ctrl)
	s := mock.NewMockSession(ctrl)
	ctx := context.Background()

	remote := Remote{
		Repo:       repo,
		Connection: conn,
		Session:    s,
	}

	conn.EXPECT().Register(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).DoAndReturn(func(_ context.Context, login string, _ string) (string, error) {
		if login == "registration_error" {
			return "", errors.New("error")
		}
		return login, nil
	}).AnyTimes()

	s.EXPECT().Save(gomock.Any()).DoAndReturn(func(session string) error {
		if session == "save_session_error" {
			return errors.New("error")
		}
		return nil
	}).AnyTimes()

	err := remote.Register(ctx, "registration_error", "password")
	assert.Error(t, err)
	err = remote.Register(ctx, "save_session_error", "password")
	assert.Error(t, err)
	err = remote.Register(ctx, "login", "password")
	assert.Nil(t, err)
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepoZip(ctrl)
	conn := mock.NewMockExtConnection(ctrl)
	s := mock.NewMockSession(ctrl)
	ctx := context.Background()

	remote := Remote{
		Repo:       repo,
		Connection: conn,
		Session:    s,
	}

	conn.EXPECT().Login(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).DoAndReturn(func(_ context.Context, login string, _ string) (string, error) {
		if login == "registration_error" {
			return "", errors.New("error")
		}
		return login, nil
	}).AnyTimes()

	s.EXPECT().Save(gomock.Any()).DoAndReturn(func(session string) error {
		if session == "save_session_error" {
			return errors.New("error")
		}
		return nil
	}).AnyTimes()

	err := remote.Login(ctx, "registration_error", "password")
	assert.Error(t, err)
	err = remote.Login(ctx, "save_session_error", "password")
	assert.Error(t, err)
	err = remote.Login(ctx, "login", "password")
	assert.Nil(t, err)
}
