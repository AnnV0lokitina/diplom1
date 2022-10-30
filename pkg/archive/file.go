package archive

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
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

// GetModTime Return the information about file.
func (f *File) GetModTime(zipName string) (time.Time, error) {
	path := filepath.Join(f.zipStorePath, zipName)
	stat, err := os.Stat(path)
	if os.IsNotExist(err) || stat.Size() == 0 {
		log.Warning("no local content")
		return time.Time{}, errors.New("no content")
	}
	return stat.ModTime(), nil
}
