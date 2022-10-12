package repo

import (
	"github.com/AnnV0lokitina/diplom1/pkg/file"
	"io"
	"path/filepath"
)

func (r *Repo) WriteFileByChunks(reader io.Reader) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	filePath := filepath.Join(r.storePath, dataFileName)
	f := file.File{
		Path: filePath,
	}
	return f.WriteByChunks(reader)
}
