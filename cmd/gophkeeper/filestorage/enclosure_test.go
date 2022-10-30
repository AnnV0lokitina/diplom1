package filestorage

import (
	"bytes"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

const testStoreDir = "test_store"
const testFileName = "test_copy_name"
const testFileContent = "pksdskjdfgla'wejgf'oowejg'powegj'pweogjap'rwojgaporegjpiwegjoiwrgjperoohgpoerhjoreierjhieo" +
	";lskdfnl;rjgkjnao;ragoaiergoiaregohiarehoijragiohraiorgijrgjripreojhphotjpehtjprhejiphrijrhjihijerhjihjij" +
	"pksoeafeofpjepsijfesfjpiewfjpiefwipjeifjppjieejipejigpjgiegwrjigijpgjipgewjkigiejgjeijigegoegoegoiewfoi[kogr"

func TestNewEnclosure(t *testing.T) {
	en := NewEnclosure(testStoreDir)
	assert.Equal(t, testStoreDir, en.storePath)
}

func TestSaveOpenRemove(t *testing.T) {
	tmpDir := os.TempDir()
	testDir, err := os.MkdirTemp(tmpDir, testStoreDir)
	require.NoError(t, err)
	reader := entity.NewTextReadCloser(testFileContent)

	enIncorrect := NewEnclosure(testDir + "1")

	err = enIncorrect.Save(testFileName, reader)
	assert.Error(t, err)

	enCorrect := NewEnclosure(testDir)
	err = enCorrect.Save(testFileName, reader)
	assert.Nil(t, err)

	_, err = enCorrect.Open(testFileName + "1")
	assert.Error(t, err)

	newReader, err := enCorrect.Open(testFileName)
	assert.Nil(t, err)
	buf := make([]byte, 16)
	data := bytes.NewBuffer([]byte{})
	for {
		n, err := newReader.Read(buf)
		if err == io.EOF {
			break
		}
		data.Write(buf[:n])
	}
	assert.Equal(t, testFileContent, data.String())

	err = enCorrect.Remove(testFileName + "1")
	assert.Error(t, err)

	err = enCorrect.Remove(testFileName)
	assert.Nil(t, err)

	_, err = enCorrect.Open(testFileName)
	assert.Error(t, err)
}
