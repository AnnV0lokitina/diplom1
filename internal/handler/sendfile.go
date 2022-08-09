package handler

import (
	pb "github.com/AnnV0lokitina/diplom1/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) SendFile(*pb.UserRequest, pb.SecureStorage_SendFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SendFile not implemented")
}
