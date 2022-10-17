package repo

import (
	"github.com/AnnV0lokitina/diplom1/internal/entity"
	"github.com/AnnV0lokitina/diplom1/pkg/file"
	zipPkg "github.com/AnnV0lokitina/diplom1/pkg/zip"
	log "github.com/sirupsen/logrus"
	"io"
	"time"
)

func (r *Repo) CreateZIP() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	log.Println(r.storePath, r.zipPath)
	return zipPkg.Pack(r.storePath, r.zipPath)
}

func (r *Repo) UnpackZIP() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return zipPkg.Unpack(r.zipPath, r.storePath)
}

func (r *Repo) GetInfo() (*entity.FileInfo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	f := file.File{
		Path: r.zipPath,
	}
	info, err := f.GetInfo()
	if err != nil {
		return &entity.FileInfo{
			UpdateTime: time.Time{},
		}, err
	}
	return &entity.FileInfo{
		UpdateTime: info.ModTime(),
	}, nil
}

func (r *Repo) ReadFileByChunks(w io.Writer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	f := file.File{
		Path: r.zipPath,
	}
	return f.ReadByChunks(w)
}

func (r *Repo) WriteFileByChunks(reader io.Reader) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	f := file.File{
		Path: r.zipPath,
	}
	return f.WriteByChunks(reader)
}
