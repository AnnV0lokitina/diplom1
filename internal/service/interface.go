package service

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	"io"
)

// DB database interface.
type DB interface {
	Close(ctx context.Context) error
	CreateUser(
		ctx context.Context,
		sessionID string,
		login string,
		passwordHash string,
	) error
	AuthUser(
		ctx context.Context,
		login string,
		passwordHash string,
	) (int, error)
	AddUserSession(ctx context.Context, user *entity.User) error
	GetUserBySessionID(ctx context.Context, activeSessionID string) (*entity.User, error)
}

// Repo file repository interface.
type Repo interface {
	GetInfo(fileName string) (*entity.FileInfo, error)
	Read(fileName string, w io.Writer) error
	Write(fileName string, reader io.Reader) error
}
