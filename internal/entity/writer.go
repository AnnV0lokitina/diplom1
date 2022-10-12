package entity

import (
	"os"
)

// Writer Store file pointer and decoder to write to file.
type Writer struct {
	file *os.File
}

// NewWriter Create new writer.
func NewWriter(filePath string) (*Writer, error) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return &Writer{
		file: file,
	}, nil
}

// WriteRecord Write record to file.
func (w *Writer) Write(fileData []byte) (int, error) {
	return w.file.Write(fileData)
}

// Close Stop work with file.
func (w *Writer) Close() error {
	return w.file.Close()
}
