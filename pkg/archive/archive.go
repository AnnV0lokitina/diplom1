package archive

import (
	"io"
	"time"
)

// Archive keep information about file.
type Archive struct {
	File
	zipName   string
	sourceDir string
}

func NewArchive(sourceDir string, zipStorePath string, zipName string) *Archive {
	return &Archive{
		File: File{
			zipStorePath: zipStorePath,
		},
		zipName:   zipName,
		sourceDir: sourceDir,
	}
}

// GetZIPModTime Return the time of last modification.
func (arch *Archive) GetZIPModTime() (time.Time, error) {
	return arch.GetModTime(arch.zipName)
}

// ReadZIPByChunks Read file by chunks.
func (arch *Archive) ReadZIPByChunks(w io.Writer) error {
	return arch.ReadByChunks(arch.zipName, w)
}

// WriteZIPByChunks Write file by chunks.
func (arch *Archive) WriteZIPByChunks(r io.Reader) error {
	return arch.WriteByChunks(arch.zipName, r)
}
