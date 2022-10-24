package filestorage

import (
	"encoding/json"
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"os"
)

var ErrorNoInfo = errors.New("no file or file empty")

// Reader Store file pointer and decoder to read file.
type Reader struct {
	file    *os.File
	decoder *json.Decoder
}

// NewReader create Reader.
func NewReader(filePath string) (*Reader, error) {
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) || stat.Size() == 0 {
		return &Reader{}, ErrorNoInfo
	}
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	return &Reader{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

// Empty Check reader empty.
func (r *Reader) Empty() bool {
	return r.file == nil
}

// ReadRecord Read record from file.
func (r *Reader) ReadRecord() (*entity.Record, error) {
	var record entity.Record
	if err := r.decoder.Decode(&record); err != nil {
		return nil, err
	}
	return &record, nil
}

// Close Stop work with file.
func (r *Reader) Close() error {
	return r.file.Close()
}
