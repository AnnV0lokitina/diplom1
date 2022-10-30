package _interface

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"io"
)

type RepoZip interface {
	CreateZIP() error
	UnpackZIP() error
	GetInfo() (*entity.FileInfo, error)
	ReadFileByChunks(w io.Writer) error
	WriteFileByChunks(reader io.Reader) error
}
