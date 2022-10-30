package _interface

import "io"

type External interface {
	Open(filePath string) (string, io.ReadCloser, error)
	Save(filePath string, reader io.ReadCloser) error
}
