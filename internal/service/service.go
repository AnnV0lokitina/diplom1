package service

import (
	"context"
	"errors"
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	labelError "github.com/AnnV0lokitina/diplom1/pkg/error"
	log "github.com/sirupsen/logrus"
	"io"
	"time"
)

// Service keep information to execute application tasks.
type Service struct {
	db   DB
	repo Repo
}

// NewService Create new Service struct.
func NewService(db DB, repo Repo) *Service {
	return &Service{
		db:   db,
		repo: repo,
	}
}

// RegisterUser Register new user.
func (s *Service) RegisterUser(ctx context.Context, login string, password string) (*entity.User, error) {
	passwordHash := entity.CreatePasswordHash(password)
	sessionID, err := entity.GenerateSessionID()
	if err != nil {
		return nil, err
	}
	err = s.db.CreateUser(ctx, sessionID, login, passwordHash)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		ActiveSessionID: sessionID,
		Login:           login,
	}
	return user, nil
}

// LoginUser Authorize user by login and password.
func (s *Service) LoginUser(ctx context.Context, login string, password string) (*entity.User, error) {
	passwordHash := entity.CreatePasswordHash(password)
	sessionID, err := entity.GenerateSessionID()
	if err != nil {
		return nil, err
	}
	userID, err := s.db.AuthUser(ctx, login, passwordHash)
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
	err = s.db.AddUserSession(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// authorizeUser Authorize user by session.
func (s *Service) authorizeUser(ctx context.Context, sessionID string) (*entity.User, error) {
	user, err := s.db.GetUserBySessionID(ctx, sessionID)
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

// RestoreFile Send new information to user.
func (s *Service) RestoreFile(
	ctx context.Context,
	sessionID string,
	time time.Time,
	w io.Writer,
) error {
	log.Info("start authorize user")
	user, err := s.authorizeUser(ctx, sessionID)
	if err != nil {
		log.Errorf("error while authorization: %s", err)
		return err
	}
	log.Info("start restore file")
	info, err := s.repo.GetInfo(user.Login)
	if err != nil {
		log.Infof("no file to restore: %s", err)
		return err
	}
	if info.UpdateTime.Before(time) {
		log.Info("stored file is old")
		return labelError.NewLabelError(labelError.TypeUpgradeRequired, errors.New("stored file is old"))
	}
	return s.repo.Read(user.Login, w)
}

// StoreFile Save new information from user.
func (s *Service) StoreFile(ctx context.Context, sessionID string, time time.Time, r io.Reader) error {
	log.Info("start authorize user")
	user, err := s.authorizeUser(ctx, sessionID)
	if err != nil {
		log.Errorf("error while authorization: %s", err)
		return err
	}
	log.Info("start save file")
	info, err := s.repo.GetInfo(user.Login)
	if err != nil {
		log.Info("no old file in store")
		return s.repo.Write(user.Login, r)
	}
	if info.UpdateTime.After(time) {
		log.Error("received file is old")
		return labelError.NewLabelError(labelError.TypeUpgradeRequired, errors.New("received file is old"))
	}
	return s.repo.Write(user.Login, r)
}
