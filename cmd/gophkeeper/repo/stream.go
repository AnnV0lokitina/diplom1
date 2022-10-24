package repo

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"io"
	"time"
)

func (r *Repo) CreateZIP() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.archive.Pack()
}

func (r *Repo) UnpackZIP() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.archive.Unpack()
}

func (r *Repo) GetInfo() (*entity.FileInfo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	info, err := r.archive.GetInfo()
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
	return r.archive.ReadByChunks(w)
}

func (r *Repo) WriteFileByChunks(reader io.Reader) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.archive.WriteByChunks(reader)
}
