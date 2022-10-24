package archive

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// ChunkSize Default chunk size
const ChunkSize = 16

// Archive keep information about file.
type Archive struct {
	path      string
	sourceDir string
}

func NewArchive(sourceDir string, path string) *Archive {
	return &Archive{
		path:      path,
		sourceDir: sourceDir,
	}
}

// GetInfo Return the information about file.
func (f *Archive) GetInfo() (os.FileInfo, error) {
	stat, err := os.Stat(f.path)
	if os.IsNotExist(err) || stat.Size() == 0 {
		log.Warning("no local content")
		return stat, errors.New("no content")
	}
	return stat, nil
}
