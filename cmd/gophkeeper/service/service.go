package service

import (
	"context"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/interface"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/remote"
)

type Service struct {
	r    _interface.Remote
	repo _interface.RepoData
}

func NewService(repo _interface.Repo, conn _interface.ExtConnection, s _interface.Session) *Service {
	return &Service{
		r: &remote.Remote{
			Repo:       repo,
			Connection: conn,
			Session:    s,
		},
	}
}

// Register a user on a remote server.
func (s *Service) Register(ctx context.Context, login string, password string) error {
	return s.r.Register(ctx, login, password)
}

// Login Authorizes a user on a remote server.
func (s *Service) Login(ctx context.Context, login string, password string) error {
	return s.r.Login(ctx, login, password)
}
