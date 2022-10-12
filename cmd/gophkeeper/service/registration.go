package service

import (
	"context"
	"os"
)

func (s *Service) Register(ctx context.Context, login string, password string) error {
	session, err := s.connection.Register(ctx, login, password)
	if err != nil {
		return err
	}
	return os.Setenv("EXT_SESSION", session)
}

func (s *Service) Login(ctx context.Context, login string, password string) error {
	session, err := s.connection.Login(ctx, login, password)
	if err != nil {
		return err
	}
	return os.Setenv("EXT_SESSION", session)
}
