package grpctest

import (
	"context"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login User authorization by login and password.
func (h *Handler) Login(_ context.Context, in *pb.UserRegRequest) (*pb.UserRegResponse, error) {
	login := in.GetLogin()

	switch login {
	case ExistingUser:
		return &pb.UserRegResponse{
			Session: CorrectSession,
		}, nil
	default:
		return &pb.UserRegResponse{}, status.Error(codes.PermissionDenied, "Permission denied")
	}
}
