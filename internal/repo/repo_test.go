package repo

import (
	"bytes"
	"errors"
	"github.com/AnnV0lokitina/diplom1/internal/repo/mock"
	labelError "github.com/AnnV0lokitina/diplom1/pkg/error"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
	"time"
)

func TestNewRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileManager := mock.NewMockFile(ctrl)

	repo := NewRepo(fileManager)
	assert.NotNil(t, repo)

}

func TestGetInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	oldFileName := "old_file_name"
	newFileName := "new_file_name"
	now := time.Now()

	fileManager := mock.NewMockFile(ctrl)

	fileManager.EXPECT().GetModTime(gomock.Eq(oldFileName)).DoAndReturn(func(_ string) (time.Time, error) {
		return time.Time{}, labelError.NewLabelError(labelError.TypeUpgradeRequired, errors.New("error"))
	}).AnyTimes()

	fileManager.EXPECT().GetModTime(gomock.Eq(newFileName)).DoAndReturn(func(_ string) (time.Time, error) {
		return now, nil
	}).AnyTimes()

	repo := NewRepo(fileManager)
	_, err := repo.GetInfo(oldFileName)
	assert.Error(t, err)

	info, err := repo.GetInfo(newFileName)
	assert.Nil(t, err)
	assert.Equal(t, now, info.UpdateTime)

}

func TestRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileManager := mock.NewMockFile(ctrl)
	withError := "with_error"
	noError := "no_error"

	fileManager.EXPECT().ReadByChunks(
		withError,
		gomock.Any(),
	).DoAndReturn(func(_ string, w io.Writer) error {
		_, err := w.Write([]byte("one"))
		require.NoError(t, err)
		return errors.New("error")
	}).AnyTimes()

	fileManager.EXPECT().ReadByChunks(
		noError,
		gomock.Any(),
	).DoAndReturn(func(_ string, w io.Writer) error {
		_, err := w.Write([]byte("one"))
		require.NoError(t, err)
		_, err = w.Write([]byte("two"))
		require.NoError(t, err)
		_, err = w.Write([]byte("three"))
		require.NoError(t, err)
		return nil
	}).AnyTimes()

	repo := NewRepo(fileManager)
	w := bytes.Buffer{}
	err := repo.Read(withError, &w)
	assert.Error(t, err)

	w.Reset()
	err = repo.Read(noError, &w)
	assert.Nil(t, err)
	assert.Equal(t, "onetwothree", w.String())
}

func TestWrite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileManager := mock.NewMockFile(ctrl)
	withError := "with_error"
	noError := "no_error"
	text := "onetwothreeonetwothreeonetwothreeonetwothreeonetwothree"
	reader := strings.NewReader(text)
	result := bytes.Buffer{}

	fileManager.EXPECT().WriteByChunks(
		withError,
		gomock.Any(),
	).DoAndReturn(func(_ string, r io.Reader) error {
		return errors.New("error")
	}).AnyTimes()

	fileManager.EXPECT().WriteByChunks(
		noError,
		gomock.Any(),
	).DoAndReturn(func(_ string, r io.Reader) error {
		buf := make([]byte, 3)
		for {
			n, err := reader.Read(buf)
			if err != nil {
				if err != io.EOF {
					if err != nil {
						return errors.New("read file chunk error")
					}
				}
				break
			}
			if n == 0 {
				return nil
			}
			_, err = result.Write(buf[0:n])
			if err != nil {
				return errors.New("send file chunk error")
			}
		}
		return nil
	}).AnyTimes()

	repo := NewRepo(fileManager)
	err := repo.Write(withError, reader)
	assert.Error(t, err)

	err = repo.Write(noError, reader)
	assert.Nil(t, err)
	assert.Equal(t, text, result.String())
}
