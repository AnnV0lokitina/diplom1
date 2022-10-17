package handler

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	pb "github.com/AnnV0lokitina/diplom1/proto"
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

// Handler structure holds dependencies for server handlers.
type Handler struct {
	pb.UnimplementedSecureStorageServer

	service Service
}

// NewHandler Create new handler struct.
func NewHandler(service Service) *Handler {
	h := &Handler{}
	h.service = service
	return h
}

//func getUserID(id uint32) (uint32, error) {
//	if id > 0 {
//		return id, nil
//	}
//	return userid.GenerateUserID()
//}
