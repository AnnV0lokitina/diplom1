package file

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// ChunkSize Default chunk size
const ChunkSize = 16

// File keep information about file.
type File struct {
	Path string
}

// GetInfo Return the information about file.
func (f *File) GetInfo() (os.FileInfo, error) {
	stat, err := os.Stat(f.Path)
	if os.IsNotExist(err) || stat.Size() == 0 {
		log.Warning("no local content")
		return stat, errors.New("no content")
	}
	return stat, nil
}
