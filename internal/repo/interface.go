package repo

import (
	"io"
	"os"
)

type File interface {
	ReadByChunks(zipName string, w io.Writer) error
	WriteByChunks(zipName string, r io.Reader) error
	GetInfo(zipName string) (os.FileInfo, error)
}
