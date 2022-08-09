package handler

import (
	pb "github.com/AnnV0lokitina/diplom1/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetFile(pb.SecureStorage_GetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
