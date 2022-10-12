package repo

import (
	"errors"
	"github.com/AnnV0lokitina/diplom1/pkg/file"
	"io"
	"os"
	"path/filepath"
)

func (r *Repo) GetInfo() (os.FileInfo, error) {
	filePath := filepath.Join(r.storePath, dataFileName)
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) || stat.Size() == 0 {
		return stat, errors.New("no content")
	}
	return stat, nil
}

func (r *Repo) ReadFileByChunks(w io.Writer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	filePath := filepath.Join(r.storePath, dataFileName)
	f := file.File{
		Path: filePath,
	}
	return f.ReadByChunks(w)

	//f, err := os.Open(filePath)
	//if err != nil {
	//	return errors.New("open file error")
	//}
	//defer f.Close()
	//
	//reader := bufio.NewReader(f)
	//buf := make([]byte, 16)
	//
	//for {
	//	n, err := reader.Read(buf)
	//	if err != nil {
	//		if err != io.EOF {
	//			if err != nil {
	//				return errors.New("read file chunk error")
	//			}
	//		}
	//		break
	//	}
	//	if n == 0 {
	//		return nil
	//	}
	//	fmt.Println("read: ", string(buf[0:n]))
	//	_, err = w.Write(buf[0:n])
	//	if err != nil {
	//		return errors.New("send file chunk error")
	//	}
	//}
	//return nil
}
