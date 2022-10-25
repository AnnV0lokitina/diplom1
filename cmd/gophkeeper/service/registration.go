package service

import (
	"context"
)

// Register a user on a remote server.
func (s *Service) Register(ctx context.Context, login string, password string) error {
	session, err := s.connection.Register(ctx, login, password)
	if err != nil {
		return err
	}
	return s.session.Save(session)
}

// Login Authorizes a user on a remote server.
func (s *Service) Login(ctx context.Context, login string, password string) error {
	session, err := s.connection.Login(ctx, login, password)
	if err != nil {
		return err
	}
	return s.session.Save(session)
}
