package handler

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	"github.com/AnnV0lokitina/diplom1/pkg/userid"
	pb "github.com/AnnV0lokitina/diplom1/proto"
)

type Service interface {
	RegisterUser(ctx context.Context, login string, password string) (*entity.User, error)
	LoginUser(ctx context.Context, login string, password string) (*entity.User, error)
}

// Handler structure holds dependencies for server handlers.
type Handler struct {
	pb.UnimplementedSecureStorageServer

	service Service
}

func NewHandler(service Service) *Handler {
	h := &Handler{}
	h.service = service
	return h
}

func getUserID(id uint32) (uint32, error) {
	if id > 0 {
		return id, nil
	}
	return userid.GenerateUserID()
}
