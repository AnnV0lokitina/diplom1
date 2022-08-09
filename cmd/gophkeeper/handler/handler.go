package handler

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo"
)

var ConfigPath string

type Handler struct {
	params Params
	repo   *repo.Repo
}

func NewHandler(p Params) (*Handler, error) {
	err := setConfig(ConfigPath, &p)
	if err != nil {
		return nil, err
	}
	r, err := repo.NewFileRepo(p.FileStorePath)
	if err != nil {
		return nil, err
	}
	return &Handler{
		params: p,
		repo:   r,
	}, nil
}

func setConfig(path string, p *Params) error {
	if path == "" {
		return nil
	}
	return p.SetFromJSON(path)
}
