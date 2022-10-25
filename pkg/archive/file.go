package archive

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// ChunkSize Default chunk size
const ChunkSize = 16

type File struct {
	zipStorePath string
}

func NewFile(zipStorePath string) *File {
	return &File{
		zipStorePath: zipStorePath,
	}
}

// GetInfo Return the information about file.
func (f *File) GetInfo(zipName string) (os.FileInfo, error) {
	path := filepath.Join(f.zipStorePath, zipName)
	stat, err := os.Stat(path)
	if os.IsNotExist(err) || stat.Size() == 0 {
		log.Warning("no local content")
		return stat, errors.New("no content")
	}
	return stat, nil
}
