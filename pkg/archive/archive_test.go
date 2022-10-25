package archive

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

const testSourceStoreDir = "test_source_store"

func createTestSourceDir(t *testing.T) string {
	tmpDir := os.TempDir()
	testDir, err := os.MkdirTemp(tmpDir, testSourceStoreDir)
	require.NoError(t, err)

	return testDir
}

func TestNewArchive(t *testing.T) {
	zipStorePath := createTestDir(t)
	zipName := "tmp.txt"
	sourceDir := createTestSourceDir(t)
	arch := NewArchive(sourceDir, zipStorePath, zipName)
	assert.Equal(t, zipStorePath, arch.zipStorePath)
	assert.Equal(t, zipName, arch.zipName)
	assert.Equal(t, sourceDir, arch.sourceDir)
}

func TestGetZIPInfo(t *testing.T) {
	text := "eokrapoerpoaewkorpweaokdaewok"
	zipStorePath := createTestDir(t)
	zipName := "test_file.txt"
	path := createTestFile(t, zipStorePath, zipName, text)

	sourceDir := createTestSourceDir(t)
	arch := NewArchive(sourceDir, zipStorePath, zipName)
	info, err := arch.GetZIPInfo()
	assert.Nil(t, err)
	assert.Equal(t, zipName, info.Name())

	removeTestFile(t, path)
}

func TestReadZIPByChunks(t *testing.T) {
	text := "eokrapoerpoaewkorpweaokdaewok;sejfcsejfcapowejfpoaewjfpoJKEFPOJEWFOPjweofjowpejfpoejwfpowejfweojg" +
		"awekjfoaiwejfiqwejfgoiqwj4efoiwjefgoqiwrejgfiqrwgjhoirwhgjiqrwhgirhgoirhgiqhrgiqwrjgijrgojqwgpojqwpg"
	zipStorePath := createTestDir(t)
	zipName := "test_file_read.txt"
	path := createTestFile(t, zipStorePath, zipName, text)

	sourceDir := createTestSourceDir(t)
	arch := NewArchive(sourceDir, zipStorePath, zipName)
	w := bytes.Buffer{}
	err := arch.ReadZIPByChunks(&w)
	assert.Nil(t, err)
	assert.Equal(t, text, w.String())

	removeTestFile(t, path)
}

func TestWriteZIPByChunks(t *testing.T) {
	text := "eokrapoerpoaewkorpweaokdaewok;sejfcsejfcapowejfpoaewjfpoJKEFPOJEWFOPjweofjowpejfpoejwfpowejfweojg" +
		"awekjfoaiwejfiqwejfgoiqwj4efoiwjefgoqiwrejgfiqrwgjhoirwhgjiqrwhgirhgoirhgiqhrgiqwrjgijrgojqwgpojqwpg"
	r := strings.NewReader(text)

	zipStorePath := createTestDir(t)
	zipName := "test_file_read.txt"

	sourceDir := createTestSourceDir(t)
	arch := NewArchive(sourceDir, zipStorePath, zipName)
	err := arch.WriteZIPByChunks(r)
	assert.Nil(t, err)

	w := bytes.Buffer{}
	err = arch.ReadZIPByChunks(&w)
	require.NoError(t, err)
	assert.Equal(t, text, w.String())
}

func TestZIPUnZIP(t *testing.T) {

}
