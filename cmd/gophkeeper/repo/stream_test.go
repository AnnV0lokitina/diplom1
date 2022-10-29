package repo

import (
	"bytes"
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestZIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := mock.NewMockFileStorageReader(ctrl)
	writer := mock.NewMockFileStorageWriter(ctrl)
	enclosure := mock.NewMockFileStorageEnclosure(ctrl)
	archive := mock.NewMockArchive(ctrl)

	record := createTestRecord()

	r := &bytes.Buffer{}
	w := &bytes.Buffer{}
	now := time.Now()

	repo := &Repo{
		record:    record,
		writer:    writer,
		enclosure: enclosure,
		archive:   archive,
		reader:    reader,
	}

	archive.EXPECT().Pack().DoAndReturn(func() error {
		return errors.New("error")
	}).Times(1)
	archive.EXPECT().Pack().DoAndReturn(func() error {
		return nil
	}).Times(1)

	err := repo.CreateZIP()
	assert.Error(t, err)
	err = repo.CreateZIP()
	assert.Nil(t, err)

	archive.EXPECT().Unpack().DoAndReturn(func() error {
		return errors.New("error")
	}).Times(1)
	archive.EXPECT().Unpack().DoAndReturn(func() error {
		return nil
	}).Times(1)

	err = repo.UnpackZIP()
	assert.Error(t, err)
	err = repo.UnpackZIP()
	assert.Nil(t, err)

	archive.EXPECT().GetZIPModTime().DoAndReturn(func() (time.Time, error) {
		return time.Time{}, errors.New("error")
	}).Times(1)
	archive.EXPECT().GetZIPModTime().DoAndReturn(func() (time.Time, error) {
		return now, nil
	}).Times(1)

	_, err = repo.GetInfo()
	assert.Error(t, err)
	fileInfo, err := repo.GetInfo()
	assert.Nil(t, err)
	assert.Equal(t, now, fileInfo.UpdateTime)

	archive.EXPECT().ReadZIPByChunks(gomock.Any()).DoAndReturn(func(_ io.Writer) error {
		return errors.New("error")
	}).Times(1)
	archive.EXPECT().ReadZIPByChunks(gomock.Any()).DoAndReturn(func(_ io.Writer) error {
		return nil
	}).Times(1)

	err = repo.ReadFileByChunks(w)
	assert.Error(t, err)
	err = repo.ReadFileByChunks(w)
	assert.Nil(t, err)

	archive.EXPECT().WriteZIPByChunks(gomock.Any()).DoAndReturn(func(_ io.Reader) error {
		return errors.New("error")
	}).Times(1)
	archive.EXPECT().WriteZIPByChunks(gomock.Any()).DoAndReturn(func(_ io.Reader) error {
		return nil
	}).Times(1)

	err = repo.WriteFileByChunks(r)
	assert.Error(t, err)
	err = repo.WriteFileByChunks(r)
	assert.Nil(t, err)
}
