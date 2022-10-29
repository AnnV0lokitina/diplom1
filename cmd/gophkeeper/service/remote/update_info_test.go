package remote

import (
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/remote/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSendInfo(t *testing.T) {
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

	// error get info
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().CreateZIP().Return(nil)
	repo.EXPECT().GetInfo().Return(nil, errors.New("error"))

	err := remote.SendInfo(ctx)
	assert.Error(t, err)

	// error get info
	s.EXPECT().Get().Return("", errors.New("error"))
	repo.EXPECT().CreateZIP().Return(nil)
	repo.EXPECT().GetInfo().Return(nil, errors.New("error"))

	err = remote.SendInfo(ctx)
	assert.Error(t, err)

	// error create zip
	s.EXPECT().Get().Return("", errors.New("error"))
	repo.EXPECT().CreateZIP().Return(errors.New("error"))

	err = remote.SendInfo(ctx)
	assert.Error(t, err)

	// error store info
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().CreateZIP().Return(nil)
	repo.EXPECT().GetInfo().DoAndReturn(func() (*entity.FileInfo, error) {
		return &entity.FileInfo{
			UpdateTime: time.Now(),
		}, nil
	})
	repo.EXPECT().ReadFileByChunks(gomock.Any()).Return(nil)
	conn.EXPECT().StoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))

	err = remote.SendInfo(ctx)
	assert.Error(t, err)

	// error store info and read file
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().CreateZIP().Return(nil)
	repo.EXPECT().GetInfo().DoAndReturn(func() (*entity.FileInfo, error) {
		return &entity.FileInfo{
			UpdateTime: time.Now(),
		}, nil
	})
	repo.EXPECT().ReadFileByChunks(gomock.Any()).Return(errors.New("error"))
	conn.EXPECT().StoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))

	err = remote.SendInfo(ctx)
	assert.Error(t, err)

	// error  read file
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().CreateZIP().Return(nil)
	repo.EXPECT().GetInfo().DoAndReturn(func() (*entity.FileInfo, error) {
		return &entity.FileInfo{
			UpdateTime: time.Now(),
		}, nil
	})
	repo.EXPECT().ReadFileByChunks(gomock.Any()).Return(errors.New("error"))
	conn.EXPECT().StoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err = remote.SendInfo(ctx)
	assert.Error(t, err)

	// no error
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().CreateZIP().Return(nil)
	repo.EXPECT().GetInfo().DoAndReturn(func() (*entity.FileInfo, error) {
		return &entity.FileInfo{
			UpdateTime: time.Now(),
		}, nil
	})
	repo.EXPECT().ReadFileByChunks(gomock.Any()).Return(nil)
	conn.EXPECT().StoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err = remote.SendInfo(ctx)
	assert.Nil(t, err)
}

func TestReceiveInfo(t *testing.T) {
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

	// error write file
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().GetInfo().DoAndReturn(func() (*entity.FileInfo, error) {
		return &entity.FileInfo{
			UpdateTime: time.Now(),
		}, nil
	})
	repo.EXPECT().WriteFileByChunks(gomock.Any()).Return(errors.New("error"))
	conn.EXPECT().RestoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err := remote.ReceiveInfo(ctx)
	assert.Error(t, err)

	// error write file
	s.EXPECT().Get().Return("session", errors.New("error"))
	repo.EXPECT().GetInfo().DoAndReturn(func() (*entity.FileInfo, error) {
		return &entity.FileInfo{
			UpdateTime: time.Now(),
		}, nil
	})
	repo.EXPECT().WriteFileByChunks(gomock.Any()).Return(errors.New("error"))
	conn.EXPECT().RestoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err = remote.ReceiveInfo(ctx)
	assert.Error(t, err)

	// error write file
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().GetInfo().Return(nil, errors.New("error"))
	repo.EXPECT().WriteFileByChunks(gomock.Any()).Return(errors.New("error"))
	conn.EXPECT().RestoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err = remote.ReceiveInfo(ctx)
	assert.Error(t, err)

	// error restore file
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().GetInfo().Return(nil, errors.New("error"))
	repo.EXPECT().WriteFileByChunks(gomock.Any()).Return(nil)
	conn.EXPECT().RestoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))

	err = remote.ReceiveInfo(ctx)
	assert.Error(t, err)

	// error unpack zip
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().GetInfo().Return(nil, errors.New("error"))
	repo.EXPECT().WriteFileByChunks(gomock.Any()).Return(nil)
	conn.EXPECT().RestoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	repo.EXPECT().UnpackZIP().Return(errors.New("error"))

	err = remote.ReceiveInfo(ctx)
	assert.Error(t, err)

	// no error
	s.EXPECT().Get().Return("session", nil)
	repo.EXPECT().GetInfo().Return(nil, errors.New("error"))
	repo.EXPECT().WriteFileByChunks(gomock.Any()).Return(nil)
	conn.EXPECT().RestoreInfo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	repo.EXPECT().UnpackZIP().Return(nil)

	err = remote.ReceiveInfo(ctx)
	assert.Nil(t, err)
}
