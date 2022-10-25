package filestorage

import (
	"os"
	"path/filepath"
)

type Session struct {
	fileName string
}

func NewSession(fileName string) *Session {
	return &Session{
		fileName: fileName,
	}
}

// Save Saves session id.
func (s *Session) Save(sessionID string) error {
	path := filepath.Join(os.TempDir(), s.fileName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(sessionID)
	return err
}

// Get Receives session id.
func (s *Session) Get() (string, error) {
	path := filepath.Join(os.TempDir(), s.fileName)
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
