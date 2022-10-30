package filestorage

import (
	"errors"
	"io"
	"os"
)

type ExternalFile struct {
}

func NewExternalFile() *ExternalFile {
	return &ExternalFile{}
}

// Save file from reader with name.
func (ef *ExternalFile) Save(filePath string, reader io.ReadCloser) error {
	defer reader.Close()
	fo, err := os.Create(filePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(fo, reader)
	if err != nil {
		return err
	}
	return nil
}

// Open file by name and return reader.
func (ef *ExternalFile) Open(filePath string) (string, io.ReadCloser, error) {
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) || stat.Size() == 0 {
		return "", nil, errors.New("no source file")
	}
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		return "", nil, err
	}
	return stat.Name(), file, nil
}
