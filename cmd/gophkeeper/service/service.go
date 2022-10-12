package service

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/external"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo"
)

type Service struct {
	repo       *repo.Repo
	connection *external.ExtConnection
}

func NewService(repo *repo.Repo, conn *external.ExtConnection) *Service {
	return &Service{
		repo:       repo,
		connection: conn,
	}
}
