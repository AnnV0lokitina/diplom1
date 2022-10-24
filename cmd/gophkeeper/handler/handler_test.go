package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHandler(t *testing.T) {
	p := Params{
		ServerAddress: "localhost:3200",
		FileStorePath: "",
		ArchiveName:   "user_archive.zip",
		DataFileName:  "data.json",
	}

	h, err := NewHandler(p)
	assert.Nil(t, err)
	assert.Equal(t, p, h.params)
}
