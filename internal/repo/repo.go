package repo

import (
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	labelError "github.com/AnnV0lokitina/diplom1/pkg/error"
	"github.com/AnnV0lokitina/diplom1/pkg/file"
	log "github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"sync"
)

// Repo keep information about file repository.
type Repo struct {
	mu        sync.Mutex
	storePath string
}

// NewRepo Create new Repo struct.
func NewRepo(storePath string) *Repo {
	return &Repo{
		storePath: storePath,
	}
}

// GetInfo get information about file.
func (r *Repo) GetInfo(fileName string) (*entity.FileInfo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	file := file.File{
		Path: filepath.Join(r.storePath, fileName),
	}
	info, err := file.GetInfo()
	if err != nil {
		return nil, labelError.NewLabelError(labelError.TypeUpgradeRequired, err)
	}
	return &entity.FileInfo{
		UpdateTime: info.ModTime(),
	}, nil
}

// Read file to stream
func (r *Repo) Read(fileName string, w io.Writer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	source := file.File{
		Path: filepath.Join(r.storePath, fileName),
	}

	log.Infof("read file by chunks path: %s", source.Path)
	return source.ReadByChunks(w)
}

// Write file from stream
func (r *Repo) Write(fileName string, reader io.Reader) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	source := file.File{
		Path: filepath.Join(r.storePath, fileName),
	}

	log.Infof("write file by chunks path: %s", source.Path)
	return source.WriteByChunks(reader)
}
