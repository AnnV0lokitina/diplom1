package service

import (
	"os"
	"path/filepath"
)

// SaveSession Saves session id.
func SaveSession(sessionID string) error {
	path := filepath.Join(os.TempDir(), "keeper_session.txt")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(sessionID)
	return err
}

// GetSession Receives session id.
func GetSession() (string, error) {
	path := filepath.Join(os.TempDir(), "keeper_session.txt")
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
