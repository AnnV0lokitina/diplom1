package file

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func (f *File) ReadByChunks(w io.Writer) error {
	file, err := os.Open(f.Path)
	if err != nil {
		return errors.New("open file error")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, chunkSize)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				if err != nil {
					return errors.New("read file chunk error")
				}
			}
			break
		}
		if n == 0 {
			return nil
		}
		fmt.Println("read: ", string(buf[0:n]))
		_, err = w.Write(buf[0:n])
		if err != nil {
			return errors.New("send file chunk error")
		}
	}
	return nil
}
