package handler

import (
	"context"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login User authorization by login and password.
func (h *Handler) Login(ctx context.Context, in *pb.UserRegRequest) (*pb.UserRegResponse, error) {
	password := in.GetPassword()
	login := in.GetLogin()

	user, err := h.service.LoginUser(ctx, login, password)
	if err != nil {
		return &pb.UserRegResponse{}, status.Error(codes.PermissionDenied, "Permission denied")
	}
	return &pb.UserRegResponse{
		Session: user.ActiveSessionID,
	}, nil
}
