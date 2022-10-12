package handler

import (
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/external"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service"
)

var ConfigPath string

type Handler struct {
	params  Params
	service *service.Service
}

func NewHandler(p Params) (*Handler, error) {
	if ConfigPath == "" {
		return nil, errors.New("no config")
	}
	err := setConfig(ConfigPath, &p)
	if err != nil {
		return nil, err
	}
	r, err := repo.NewFileRepo(p.FileStorePath)
	if err != nil {
		return nil, err
	}
	c := external.NewExtConnection(p.ServerAddress, p.FileStorePath)
	s := service.NewService(r, c)
	return &Handler{
		params:  p,
		service: s,
	}, nil
}

func setConfig(path string, p *Params) error {
	if path == "" {
		return nil
	}
	return p.SetFromJSON(path)
}
