package filestorage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

const testSessionFileName = "test_session_file_name.txt"

func TestSession(t *testing.T) {
	s := NewSession(testSessionFileName)
	_, err := s.Get()
	assert.Error(t, err)

	sessionID := "1234"
	err = s.Save(sessionID)
	assert.Nil(t, err)

	gotSessionID, err := s.Get()
	assert.Nil(t, err)
	assert.Equal(t, sessionID, gotSessionID)

	path := filepath.Join(os.TempDir(), testSessionFileName)
	os.Remove(path)
}
