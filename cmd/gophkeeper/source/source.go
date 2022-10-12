package source

import (
	"bufio"
	"errors"
	"os"
)

func CreateSourceReader(filePath string) (*bufio.Reader, error) {
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) || stat.Size() == 0 {
		return nil, errors.New("no source file")
	}
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(file), nil
}
