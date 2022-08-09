package service

import (
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	labelError "github.com/AnnV0lokitina/diplom1/pkg/error"
	log "github.com/sirupsen/logrus"
)

type Repo interface {
	Close(ctx context.Context) error
	CreateUser(
		ctx context.Context,
		sessionID string,
		login string,
		passwordHash string,
	) error
	AuthUser(
		ctx context.Context,
		login string,
		passwordHash string,
	) (int, error)
	AddUserSession(ctx context.Context, user *entity.User) error
	GetUserBySessionID(ctx context.Context, activeSessionID string) (*entity.User, error)
}

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) RegisterUser(ctx context.Context, login string, password string) (*entity.User, error) {
	passwordHash := entity.CreatePasswordHash(password)
	sessionID, err := entity.GenerateSessionID()
	if err != nil {
		return nil, err
	}
	err = s.repo.CreateUser(ctx, sessionID, login, passwordHash)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		ActiveSessionID: sessionID,
		Login:           login,
	}
	return user, nil
}

func (s *Service) LoginUser(ctx context.Context, login string, password string) (*entity.User, error) {
	passwordHash := entity.CreatePasswordHash(password)
	sessionID, err := entity.GenerateSessionID()
	if err != nil {
		return nil, err
	}
	userID, err := s.repo.AuthUser(ctx, login, passwordHash)
	if err != nil {
		var labelErr *labelError.LabelError
		if errors.As(err, &labelErr) && labelErr.Label == labelError.TypeNotFound {
			log.Info("user not found")
			return nil, labelError.NewLabelError(labelError.TypeUnauthorized, errors.New("user unauthorized"))
		}
		return nil, err
	}
	user := &entity.User{
		ID:              userID,
		Login:           login,
		ActiveSessionID: sessionID,
	}
	err = s.repo.AddUserSession(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) authorizeUser(ctx context.Context, sessionID string) (*entity.User, error) {
	user, err := s.repo.GetUserBySessionID(ctx, sessionID)
	if err != nil {
		var labelErr *labelError.LabelError
		if errors.As(err, &labelErr) && labelErr.Label == labelError.TypeNotFound {
			log.Info("user not found")
			return nil, labelError.NewLabelError(labelError.TypeUnauthorized, errors.New("user unauthorized"))
		}
		return nil, err
	}
	return user, nil
}
