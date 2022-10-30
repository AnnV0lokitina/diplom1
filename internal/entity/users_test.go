package entity

import (
	"crypto/md5"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func createMD5() string {
	bytePassword := []byte("password")
	idByte := md5.Sum(bytePassword)
	return fmt.Sprintf("%x", idByte)
}

func TestCreatePasswordHash(t *testing.T) {
	hash := CreatePasswordHash("password")
	assert.Equal(t, createMD5(), hash)
}

func TestGenerateSessionID(t *testing.T) {
	s1, err := GenerateSessionID()
	require.NoError(t, err)
	s2, err := GenerateSessionID()
	require.NoError(t, err)
	assert.NotEqual(t, s1, s2)
}
