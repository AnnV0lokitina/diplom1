package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	"github.com/AnnV0lokitina/diplom1/internal/service/mock"
	labelError "github.com/AnnV0lokitina/diplom1/pkg/error"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	db := mock.NewMockDB(ctrl)
	repo := mock.NewMockRepo(ctrl)

	service := NewService(db, repo)

	db.EXPECT().CreateUser(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).DoAndReturn(func(_ context.Context, _ string, _ string, _ string) error {
		return errors.New("error")
	}).Times(1)

	db.EXPECT().CreateUser(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).DoAndReturn(func(_ context.Context, _ string, login string, _ string) error {
		return nil
	}).Times(1)

	_, err := service.RegisterUser(ctx, "login", "password")
	assert.Error(t, err)
	user, err := service.RegisterUser(ctx, "login", "password")
	assert.Nil(t, err)
	assert.Equal(t, "login", user.Login)
}

func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	errUnauth := labelError.NewLabelError(labelError.TypeUnauthorized, errors.New("user unauthorized"))

	db := mock.NewMockDB(ctrl)
	repo := mock.NewMockRepo(ctrl)

	service := NewService(db, repo)

	db.EXPECT().AuthUser(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).DoAndReturn(func(_ context.Context, login string, _ string) (int, error) {
		if login == "unauthorized" {
			return 0, errUnauth
		}
		if login == "with_error" {
			return 0, errors.New("error")
		}
		return 1, nil
	}).Times(4)

	db.EXPECT().AddUserSession(
		gomock.Any(),
		gomock.Any(),
	).DoAndReturn(func(_ context.Context, user *entity.User) error {
		if user.Login == "err_add_session" {
			return errors.New("error")
		}
		return nil
	}).Times(2)

	_, err := service.LoginUser(ctx, "unauthorized", "password")
	assert.Equal(t, errUnauth, err)

	_, err = service.LoginUser(ctx, "with_error", "password")
	assert.Error(t, err)

	_, err = service.LoginUser(ctx, "err_add_session", "password")
	assert.Error(t, err)

	user, err := service.LoginUser(ctx, "login", "password")
	assert.Nil(t, err)
	assert.Equal(t, "login", user.Login)
}

func TestAuthorizeUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	errUnauth := labelError.NewLabelError(labelError.TypeUnauthorized, errors.New("user unauthorized"))

	db := mock.NewMockDB(ctrl)
	repo := mock.NewMockRepo(ctrl)

	service := NewService(db, repo)

	db.EXPECT().GetUserBySessionID(
		gomock.Any(),
		gomock.Any(),
	).DoAndReturn(func(_ context.Context, session string) (*entity.User, error) {
		if session == "unauthorized" {
			return nil, errUnauth
		}
		if session == "with_error" {
			return nil, errors.New("error")
		}
		return &entity.User{
			ID:              1,
			Login:           "login",
			ActiveSessionID: session,
		}, nil
	}).Times(3)

	_, err := service.authorizeUser(ctx, "unauthorized")
	assert.Equal(t, errUnauth, err)
	_, err = service.authorizeUser(ctx, "with_error")
	assert.Error(t, err)
	user, err := service.authorizeUser(ctx, "session")
	assert.Nil(t, err)
	assert.Equal(t, "session", user.ActiveSessionID)
}

func TestRestoreFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	errUnauth := labelError.NewLabelError(labelError.TypeUnauthorized, errors.New("user unauthorized"))
	errUpgrade := labelError.NewLabelError(labelError.TypeUpgradeRequired, errors.New("stored file is old"))
	now := time.Now()
	yesterday := now.Add(time.Hour * -24)
	twoDaysAgo := now.Add(time.Hour * -48)
	w := &bytes.Buffer{}

	db := mock.NewMockDB(ctrl)
	repo := mock.NewMockRepo(ctrl)

	service := NewService(db, repo)

	db.EXPECT().GetUserBySessionID(
		gomock.Any(),
		gomock.Any(),
	).DoAndReturn(func(_ context.Context, session string) (*entity.User, error) {
		if session == "unauthorized" {
			return nil, errUnauth
		}
		return &entity.User{
			ID:              1,
			Login:           session,
			ActiveSessionID: session,
		}, nil
	}).AnyTimes()

	repo.EXPECT().GetInfo(
		gomock.Any(),
	).DoAndReturn(func(user string) (*entity.FileInfo, error) {
		if user == "get_info_error" {
			return nil, errors.New("get_info_error")
		}
		if user == "old_file" {
			return &entity.FileInfo{
				UpdateTime: now,
			}, nil
		}
		return &entity.FileInfo{
			UpdateTime: twoDaysAgo,
		}, nil
	}).AnyTimes()

	repo.EXPECT().Read(gomock.Any(), gomock.Any()).DoAndReturn(func(user string, _ io.Writer) error {
		if user == "write_error" {
			return errors.New("write_error")
		}
		return nil
	}).AnyTimes()

	err := service.RestoreFile(ctx, "unauthorized", yesterday, w)
	assert.Equal(t, errUnauth, err)
	err = service.RestoreFile(ctx, "get_info_error", yesterday, w)
	assert.Error(t, err)
	err = service.RestoreFile(ctx, "old_file", yesterday, w)
	assert.Equal(t, errUpgrade, err)
	err = service.RestoreFile(ctx, "write_error", yesterday, w)
	assert.Error(t, err)
	err = service.RestoreFile(ctx, "session", yesterday, w)
	assert.Nil(t, err)
}

func TestStoreFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	errUnauth := labelError.NewLabelError(labelError.TypeUnauthorized, errors.New("user unauthorized"))
	errUpgrade := labelError.NewLabelError(labelError.TypeUpgradeRequired, errors.New("stored file is old"))
	now := time.Now()
	yesterday := now.Add(time.Hour * -24)
	twoDaysAgo := now.Add(time.Hour * -48)
	w := &bytes.Buffer{}

	db := mock.NewMockDB(ctrl)
	repo := mock.NewMockRepo(ctrl)

	service := NewService(db, repo)

	db.EXPECT().GetUserBySessionID(
		gomock.Any(),
		gomock.Any(),
	).DoAndReturn(func(_ context.Context, session string) (*entity.User, error) {
		if session == "unauthorized" {
			return nil, errUnauth
		}
		return &entity.User{
			ID:              1,
			Login:           session,
			ActiveSessionID: session,
		}, nil
	}).AnyTimes()

	repo.EXPECT().GetInfo(
		gomock.Any(),
	).DoAndReturn(func(user string) (*entity.FileInfo, error) {
		if user == "get_info_error" {
			return nil, errors.New("get_info_error")
		}
		if user == "old_file" {
			return &entity.FileInfo{
				UpdateTime: now,
			}, nil
		}
		return &entity.FileInfo{
			UpdateTime: twoDaysAgo,
		}, nil
	}).AnyTimes()

	repo.EXPECT().Read(gomock.Any(), gomock.Any()).DoAndReturn(func(user string, _ io.Writer) error {
		if user == "write_error" {
			return errors.New("write_error")
		}
		return nil
	}).AnyTimes()

	err := service.RestoreFile(ctx, "unauthorized", yesterday, w)
	assert.Equal(t, errUnauth, err)
	err = service.RestoreFile(ctx, "get_info_error", yesterday, w)
	assert.Error(t, err)
	err = service.RestoreFile(ctx, "old_file", yesterday, w)
	assert.Equal(t, errUpgrade, err)
	err = service.RestoreFile(ctx, "write_error", yesterday, w)
	assert.Error(t, err)
	err = service.RestoreFile(ctx, "session", yesterday, w)
	assert.Nil(t, err)
}
