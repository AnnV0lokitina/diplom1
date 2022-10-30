package filestorage

import (
	"bytes"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"path/filepath"
	"testing"
)

const testExtStoreDir = "test_ext_store"
const testExtFileName = "test_ext_copy_name"
const testExtFileContent = "pksdskjdfgla'wejgf'oowejg'powegj'pweogjap'rwojgaporegjpiwegjoiwrgjperoohgpoerhjoreierjhieo" +
	";lskdfnl;rjgkjnao;ragoaiergoiaregohiarehoijragiohraiorgijrgjripreojhphotjpehtjprhejiphrijrhjihijerhjihjij" +
	"pksoeafeofpjepsijfesfjpiewfjpiefwipjeifjppjieejipejigpjgiegwrjigijpgjipgewjkigiejgjeijigegoegoegoiewfoi[kogr"

func TestExtSaveOpen(t *testing.T) {
	tmpDir := os.TempDir()
	testDir, err := os.MkdirTemp(tmpDir, testExtStoreDir)
	require.NoError(t, err)
	reader := entity.NewTextReadCloser(testExtFileContent)

	pathIncorrect := filepath.Join(testDir+"1", testExtFileName)
	pathCorrect := filepath.Join(testDir, testExtFileName)
	pathIncorrectFile := filepath.Join(testDir, testExtFileName+"1")

	ext := NewExternalFile()

	err = ext.Save(pathIncorrect, reader)
	assert.Error(t, err)

	err = ext.Save(pathCorrect, reader)
	assert.Nil(t, err)

	_, _, err = ext.Open(pathIncorrectFile)
	assert.Error(t, err)

	name, newReader, err := ext.Open(pathCorrect)
	assert.Nil(t, err)
	assert.Equal(t, testExtFileName, name)
	buf := make([]byte, 16)
	data := bytes.NewBuffer([]byte{})
	for {
		n, err := newReader.Read(buf)
		if err == io.EOF {
			break
		}
		data.Write(buf[:n])
	}
	assert.Equal(t, testExtFileContent, data.String())
}
