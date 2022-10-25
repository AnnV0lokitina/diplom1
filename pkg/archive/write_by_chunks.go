package archive

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

// WriteByChunks Write file by chunks.
func (f *File) WriteByChunks(zipName string, r io.Reader) error {
	path := filepath.Join(f.zipStorePath, zipName)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	b := make([]byte, ChunkSize)
	for {
		n, err := r.Read(b)
		if err == io.EOF || n == 0 {
			break
		}
		n, err = file.Write(b[:n])
		if err != nil {
			log.Error("write file chunk to file error")
			return err
		}
	}
	return nil
}
