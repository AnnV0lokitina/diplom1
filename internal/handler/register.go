package handler

import (
	"context"
	"errors"
	labelError "github.com/AnnV0lokitina/diplom1/pkg/error"
	pb "github.com/AnnV0lokitina/diplom1/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Register(ctx context.Context, in *pb.UserRegRequest) (*pb.UserRegResponse, error) {
	password := in.GetPassword()
	login := in.GetLogin()

	user, err := h.service.RegisterUser(ctx, login, password)
	if err != nil {
		var labelErr *labelError.LabelError
		if errors.As(err, &labelErr) && labelErr.Label == labelError.TypeConflict {
			log.Info("login existed")
			return &pb.UserRegResponse{}, status.Error(codes.AlreadyExists, "Already exists")
		}
		return &pb.UserRegResponse{}, status.Error(codes.PermissionDenied, "Permission denied")
	}
	return &pb.UserRegResponse{
		Session: user.ActiveSessionID,
	}, nil
}
