package handler

import (
	"errors"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/external"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service"
	"os"
	"path/filepath"
)

// ConfigPath Path to the configuration file
var ConfigPath string

// Handler Contains information about the user command handler
type Handler struct {
	params  Params
	service *service.Service
}

// NewHandler Creates a new structure Handler
func NewHandler(p Params) (*Handler, error) {
	if ConfigPath == "" {
		return nil, errors.New("no config")
	}
	err := setConfig(ConfigPath, &p)
	if err != nil {
		return nil, err
	}
	zipPath := filepath.Join(os.TempDir(), p.ArchiveName)
	r, err := repo.NewFileRepo(p.FileStorePath, zipPath)
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

// setConfig Sets the configuration from a file.
func setConfig(path string, p *Params) error {
	if path == "" {
		return nil
	}
	return p.SetFromJSON(path)
}
