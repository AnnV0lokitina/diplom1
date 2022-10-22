package grpctest

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"strings"
)

func (h *Handler) ReadTestFile(w io.Writer) error {
	file := strings.NewReader(TestFileContent)
	reader := bufio.NewReader(file)
	buf := make([]byte, ChunkSize)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				if err != nil {
					log.Error("read file chunk error")
					return errors.New("read file chunk error")
				}
			}
			break
		}
		if n == 0 {
			return nil
		}
		_, err = w.Write(buf[0:n])
		if err != nil {
			log.Error("send file chunk to writer error")
			return errors.New("send file chunk error")
		}
	}
	return nil
}
