package entity

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"time"
)

const (
	sessionIDLength = 16
	TTL             = 10 * time.Minute
)

// User keep information about user.
type User struct {
	ID              int
	Login           string
	ActiveSessionID string
}

// CreatePasswordHash  Encrypts the password.
func CreatePasswordHash(password string) string {
	bytePassword := []byte(password)
	idByte := md5.Sum(bytePassword)
	return fmt.Sprintf("%x", idByte)
}

// GenerateSessionID Creates a session id.
func GenerateSessionID() (string, error) {
	b := make([]byte, sessionIDLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
