package filestorage

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

type Enclosure struct {
	storePath string
}

func NewEnclosure(storePath string) *Enclosure {
	return &Enclosure{
		storePath: storePath,
	}
}

// Save file from reader with name.
func (en *Enclosure) Save(fileName string, reader io.Reader) error {
	outFilePath := filepath.Join(en.storePath, fileName)
	fo, err := os.Create(outFilePath)
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
func (en *Enclosure) Open(fileName string) (io.Reader, error) {
	inFilePath := filepath.Join(en.storePath, fileName)
	fileReader, err := os.OpenFile(inFilePath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(fileReader), nil
}

// Remove file fy name.
func (en *Enclosure) Remove(fileName string) error {
	filePath := filepath.Join(en.storePath, fileName)
	return os.Remove(filePath)
}
