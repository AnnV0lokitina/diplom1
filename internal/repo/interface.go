package repo

import (
	"io"
	"time"
)

type File interface {
	ReadByChunks(zipName string, w io.Writer) error
	WriteByChunks(zipName string, r io.Reader) error
	GetModTime(zipName string) (time.Time, error)
}
