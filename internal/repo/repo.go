package repo

import (
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	labelError "github.com/AnnV0lokitina/diplom1/pkg/error"
	log "github.com/sirupsen/logrus"
	"io"
	"sync"
)

// Repo keep information about file repository.
type Repo struct {
	mu   sync.Mutex
	file File
}

// NewRepo Create new Repo struct.
func NewRepo(f File) *Repo {
	return &Repo{
		file: f,
	}
}

// GetInfo get information about file.
func (r *Repo) GetInfo(fileName string) (*entity.FileInfo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	modTime, err := r.file.GetModTime(fileName)
	if err != nil {
		return nil, labelError.NewLabelError(labelError.TypeUpgradeRequired, err)
	}
	return &entity.FileInfo{
		UpdateTime: modTime,
	}, nil
}

// Read file to stream
func (r *Repo) Read(fileName string, w io.Writer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	log.Infof("read file by chunks: %s", fileName)
	return r.file.ReadByChunks(fileName, w)
}

// Write file from stream
func (r *Repo) Write(fileName string, reader io.Reader) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	log.Infof("write file by chunks: %s", fileName)
	return r.file.WriteByChunks(fileName, reader)
}
