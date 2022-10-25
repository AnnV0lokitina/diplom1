package archive

import (
	"io"
	"os"
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

// GetZIPInfo Return the information about file.
func (arch *Archive) GetZIPInfo() (os.FileInfo, error) {
	return arch.GetInfo(arch.zipName)
}

// ReadZIPByChunks Read file by chunks.
func (arch *Archive) ReadZIPByChunks(w io.Writer) error {
	return arch.ReadByChunks(arch.zipName, w)
}

// WriteZIPByChunks Write file by chunks.
func (arch *Archive) WriteZIPByChunks(r io.Reader) error {
	return arch.WriteByChunks(arch.zipName, r)
}
