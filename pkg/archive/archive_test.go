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
	modTime, err := arch.GetZIPModTime()
	assert.Nil(t, err)
	stat, err := os.Stat(filepath.Join(arch.zipStorePath, arch.zipName))
	assert.Nil(t, err)
	assert.Equal(t, stat.ModTime(), modTime)

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

	path := filepath.Join(zipStorePath, zipName)
	removeTestFile(t, path)
}

func TestPackUnpack(t *testing.T) {
	text := "eokrapoerpoaewkorpweaokdaewok;sejfcsejfcapowejfpoaewjfpoJKEFPOJEWFOPjweofjowpejfpoejwfpowejfweojg" +
		"awekjfoaiwejfiqwejfgoiqwj4efoiwjefgoqiwrejgfiqrwgjhoirwhgjiqrwhgirhgoirhgiqhrgiqwrjgijrgojqwgpojqwpg"

	sourceDir := createTestSourceDir(t)
	zipStorePath := createTestDir(t)
	zipName := "test_zip.zip"
	arch := NewArchive(sourceDir, zipStorePath, zipName)

	createTestFile(t, sourceDir, "file1.txt", text)
	createTestFile(t, sourceDir, "file2.txt", text)
	createTestFile(t, sourceDir, "file3.txt", text)
	createTestFile(t, sourceDir, "file4.txt", text)

	err := arch.Pack()
	assert.Nil(t, err)

	modTime1, err := arch.GetZIPModTime()
	assert.Nil(t, err)
	stat1, err := os.Stat(filepath.Join(arch.zipStorePath, arch.zipName))
	assert.Nil(t, err)

	removeTestFile(t, filepath.Join(sourceDir, "file1.txt"))
	removeTestFile(t, filepath.Join(sourceDir, "file2.txt"))
	removeTestFile(t, filepath.Join(sourceDir, "file3.txt"))
	removeTestFile(t, filepath.Join(sourceDir, "file4.txt"))

	err = arch.Unpack()
	assert.Nil(t, err)

	removeTestFile(t, filepath.Join(zipStorePath, zipName))

	time.Sleep(time.Second * 1)

	err = arch.Pack()
	assert.Nil(t, err)

	modTime2, err := arch.GetZIPModTime()
	assert.Nil(t, err)
	stat2, err := os.Stat(filepath.Join(arch.zipStorePath, arch.zipName))
	assert.Nil(t, err)

	assert.Equal(t, true, modTime1.Before(modTime2))
	assert.Equal(t, stat1.Size(), stat2.Size())

	removeTestFile(t, filepath.Join(sourceDir, "file1.txt"))
	removeTestFile(t, filepath.Join(sourceDir, "file2.txt"))
	removeTestFile(t, filepath.Join(sourceDir, "file3.txt"))
	removeTestFile(t, filepath.Join(sourceDir, "file4.txt"))
}
