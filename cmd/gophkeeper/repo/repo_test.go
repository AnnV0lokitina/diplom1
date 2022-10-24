package repo

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFileRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := mock.NewMockFileStorageReader(ctrl)
	writer := mock.NewMockFileStorageWriter(ctrl)
	enclosure := mock.NewMockFileStorageEnclosure(ctrl)
	archive := mock.NewMockArchive(ctrl)

	reader.EXPECT().Empty().DoAndReturn(func() bool {
		return true
	}).Times(1)
	//reader.EXPECT().Close().DoAndReturn(func() error {
	//	return nil
	//}).Times(1)

	writer.EXPECT().Close().DoAndReturn(func() error {
		return nil
	}).Times(1)

	r, err := NewFileRepo(reader, writer, enclosure, archive)
	assert.Nil(t, err)
	defer r.Close()

	assert.Equal(t, archive, r.archive)
	assert.Equal(t, writer, r.writer)
	assert.Equal(t, enclosure, r.enclosure)
}
