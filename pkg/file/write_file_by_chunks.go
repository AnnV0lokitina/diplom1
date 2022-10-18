package file

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

// WriteByChunks Write file by chunks.
func (f *File) WriteByChunks(r io.Reader) error {
	file, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	b := make([]byte, ChunkSize)
	for {
		//log.Println("read start")
		n, err := r.Read(b)
		//log.Printf("n = %v err = %v b = %v\n", n, err, string(b))
		//log.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF || n == 0 {
			//log.Println("eof")
			break
		}
		n, err = file.Write(b[:n])
		//log.Println(n)
		if err != nil {
			log.Error("write file chunk to file error")
			return err
		}
	}
	return nil
}
