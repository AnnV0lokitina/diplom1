package filestorage

import (
	"encoding/json"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

const testConfigFileName = "test_config_file_name.json"

func TestSetParamsFromJSON(t *testing.T) {
	tmpDir := os.TempDir()
	testDir, err := os.MkdirTemp(tmpDir, testStoreDir)
	require.NoError(t, err)
	filePath := filepath.Join(testDir, testConfigFileName)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	require.NoError(t, err)

	params := entity.Params{
		ServerAddress: "localhost:3200",
		FileStorePath: "456",
		ArchiveName:   "user_archive.zip",
		DataFileName:  "data.json",
		Session:       "123",
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(params)
	file.Close()

	params1 := entity.Params{}
	err = SetParamsFromJSON(filePath, &params1)
	assert.Nil(t, err)
	assert.Equal(t, params, params1)

	err = SetParamsFromJSON(filePath+"1", &params1)
	assert.Error(t, err)
}
