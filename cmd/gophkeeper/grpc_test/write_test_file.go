package grpctest

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
)

// WriteTestFile Write file by chunks.
func (h *Handler) WriteTestFile(r io.Reader) error {
	file := bytes.Buffer{}
	b := make([]byte, ChunkSize)
	for {
		n, err := r.Read(b)
		if err == io.EOF || n == 0 {
			TestFileContent = file.String()
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
