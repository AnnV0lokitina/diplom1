package handler

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHandler(t *testing.T) {
	p := entity.Params{
		ServerAddress: "localhost:3200",
		FileStorePath: "",
		ArchiveName:   "user_archive.zip",
		DataFileName:  "data.json",
	}

	h, err := NewHandler(p)
	assert.Nil(t, err)
	assert.Equal(t, p, h.params)
}
