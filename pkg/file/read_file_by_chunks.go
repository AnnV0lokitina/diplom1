package file

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

// ReadByChunks Read file by chunks.
func (f *File) ReadByChunks(w io.Writer) error {
	file, err := os.Open(f.Path)
	if err != nil {
		return errors.New("open file error")
	}
	defer file.Close()

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
		//log.Info("read by chunks: ", string(buf[0:n]))
		_, err = w.Write(buf[0:n])
		if err != nil {
			log.Error("send file chunk to writer error")
			return errors.New("send file chunk error")
		}
	}
	return nil
}
