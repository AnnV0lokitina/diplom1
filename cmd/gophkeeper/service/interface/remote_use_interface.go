package _interface

import (
	"context"
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

type ExtConnection interface {
	Register(ctx context.Context, login string, password string) (string, error)
	Login(ctx context.Context, login string, password string) (string, error)
	StoreInfo(ctx context.Context, session string, reader io.Reader, info *entity.FileInfo) error
	RestoreInfo(ctx context.Context, session string, w io.Writer, fileInfo *entity.FileInfo) error
}

type Session interface {
	Save(sessionID string) error
	Get() (string, error)
}
