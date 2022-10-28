package archive

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const testStoreDir = "test_store"

func createTestDir(t *testing.T) string {
	tmpDir := os.TempDir()
	testDir, err := os.MkdirTemp(tmpDir, testStoreDir)
	require.NoError(t, err)

	return testDir
}

func createTestFile(t *testing.T, testDir string, name string, text string) string {
	path := filepath.Join(testDir, name)
	file, err := os.Create(path)
	require.NoError(t, err)
	_, err = file.WriteString(text)
	require.NoError(t, err)
	err = file.Close()
	require.NoError(t, err)

	return path
}

func removeTestFile(t *testing.T, path string) {
	err := os.Remove(path)
	require.NoError(t, err)
}

func TestNewFile(t *testing.T) {
	testDir := createTestDir(t)
	f := NewFile(testDir)
	assert.Equal(t, testDir, f.zipStorePath)
}

func TestGetInfo(t *testing.T) {
	text := "eokrapoerpoaewkorpweaokdaewok"
	testDir := createTestDir(t)
	path := createTestFile(t, testDir, "test_file.txt", text)

	f := NewFile(testDir)
	modTime, err := f.GetModTime("test_file.txt")
	assert.Nil(t, err)
	assert.IsType(t, time.Time{}, modTime)

	removeTestFile(t, path)
}

func TestReadByChunks(t *testing.T) {
	text := "eokrapoerpoaewkorpweaokdaewok;sejfcsejfcapowejfpoaewjfpoJKEFPOJEWFOPjweofjowpejfpoejwfpowejfweojg" +
		"awekjfoaiwejfiqwejfgoiqwj4efoiwjefgoqiwrejgfiqrwgjhoirwhgjiqrwhgirhgoirhgiqhrgiqwrjgijrgojqwgpojqwpg"
	testDir := createTestDir(t)
	path := createTestFile(t, testDir, "test_file_read.txt", text)

	f := NewFile(testDir)
	w := bytes.Buffer{}
	err := f.ReadByChunks("test_file_read.txt", &w)
	assert.Nil(t, err)
	assert.Equal(t, text, w.String())

	removeTestFile(t, path)
}

func TestWriteByChunks(t *testing.T) {
	text := "eokrapoerpoaewkorpweaokdaewok;sejfcsejfcapowejfpoaewjfpoJKEFPOJEWFOPjweofjowpejfpoejwfpowejfweojg" +
		"awekjfoaiwejfiqwejfgoiqwj4efoiwjefgoqiwrejgfiqrwgjhoirwhgjiqrwhgirhgoirhgiqhrgiqwrjgijrgojqwgpojqwpg"
	testDir := createTestDir(t)
	r := strings.NewReader(text)

	f := NewFile(testDir)
	err := f.WriteByChunks("test_file_write.txt", r)
	assert.Nil(t, err)

	w := bytes.Buffer{}
	err = f.ReadByChunks("test_file_write.txt", &w)
	require.NoError(t, err)
	assert.Equal(t, text, w.String())
}
