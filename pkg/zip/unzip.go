package zip

import (
	"archive/zip"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

// Unpack Get files from source archive to destination folder.
func Unpack(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		log.Errorf("unzip: cannot open reader %s", source)
		return err
	}
	defer reader.Close()

	destination, err = filepath.Abs(destination)
	if err != nil {
		log.Error("unzip: cannot get path")
		return err
	}

	for _, f := range reader.File {
		err := unzipFile(f, destination)
		if err != nil {
			log.Errorf("unzip: cannot unzip file %s", f.Name)
			return err
		}
	}

	return nil
}

func unzipFile(f *zip.File, destination string) error {
	filePath := filepath.Join(destination, f.Name)

	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		log.Errorf("unzip: file: cannot unzip file %s", filePath)
		return err
	}
	defer destinationFile.Close()

	zippedFile, err := f.Open()
	if err != nil {
		log.Error("unzip: file: cannot create zipped file")
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		log.Error("unzip: file: cannot copy zip to file")
		return err
	}
	return nil
}
