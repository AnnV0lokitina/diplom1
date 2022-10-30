package handler

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/external"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/filestorage"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service"
	"github.com/AnnV0lokitina/diplom1/pkg/archive"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// ConfigPath Path to the configuration file
var ConfigPath string

// Handler Contains information about the user command handler
type Handler struct {
	params  entity.Params
	service Service
}

// NewHandler Creates a new structure Handler
func NewHandler(p entity.Params) (*Handler, error) {
	if ConfigPath != "" {
		log.Info("load stored params")
		err := filestorage.SetParamsFromJSON(ConfigPath, &p)
		if err != nil {
			log.Errorf("load stored params error %s", err)
			return nil, err
		}
	}
	filePath := filepath.Join(p.FileStorePath, p.DataFileName)
	enclosure := filestorage.NewEnclosure(p.FileStorePath)
	log.Info("create enclosure manager")
	archive := archive.NewArchive(p.FileStorePath, os.TempDir(), p.ArchiveName)
	log.Info("create archive manager")
	writer, err := filestorage.NewWriter(filePath)
	if err != nil {
		log.Errorf("error when creating data writer manager %s", err)
		return nil, err
	}
	log.Info("create data writer manager")
	reader, err := filestorage.NewReader(filePath)
	if err != nil && err != filestorage.ErrorNoInfo {
		log.Errorf("error when creating data reader manager %s", err)
		return nil, err
	}
	log.Info("create data reader manager")
	r, err := repo.NewFileRepo(reader, writer, enclosure, archive)
	if err != nil {
		log.Errorf("create repository error %s", err)
		return nil, err
	}
	log.Info("create repository")
	c := external.NewExtConnection(p.ServerAddress, p.FileStorePath)
	log.Info("create external storage manager")
	sessStorage := filestorage.NewSession(p.Session)
	ext := filestorage.NewExternalFile()
	s := service.NewService(r, c, sessStorage, ext)
	log.Info("create service")
	return &Handler{
		params:  p,
		service: s,
	}, nil
}
