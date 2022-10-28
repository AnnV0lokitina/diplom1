package handler

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	"io"
	"time"
)

// Service declare Service interface
type Service interface {
	RegisterUser(ctx context.Context, login string, password string) (*entity.User, error)
	LoginUser(ctx context.Context, login string, password string) (*entity.User, error)
	StoreFile(ctx context.Context, sessionID string, time time.Time, r io.Reader) error
	RestoreFile(ctx context.Context, sessionID string, time time.Time, w io.Writer) error
}
