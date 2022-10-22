package grpctest

import (
	"context"
	"fmt"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register a user with a username and password.
func (h *Handler) Register(_ context.Context, in *pb.UserRegRequest) (*pb.UserRegResponse, error) {
	login := in.GetLogin()
	fmt.Println(login)

	switch login {
	case NewUser:
		return &pb.UserRegResponse{
			Session: CorrectSession,
		}, nil
	case ExistingUser:
		return &pb.UserRegResponse{}, status.Error(codes.AlreadyExists, "Already exists")
	default:
		return &pb.UserRegResponse{}, status.Error(codes.PermissionDenied, "Permission denied")
	}
}
