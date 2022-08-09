package handler

import (
	"context"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Register(context.Context, *pb.UserRequest) (*pb.TextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
