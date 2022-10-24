package archive

import (
	"archive/zip"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

func (f *Archive) Pack() error {
	log.Println("Pack", f.sourceDir, f.path)
	return packZIP(f.sourceDir, f.path)
}

// packZIP files from source folder to target archive.
func packZIP(source, target string) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name = info.Name()
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}
