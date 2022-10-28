package repo

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"io"
	"time"
)

type FileStorageReader interface {
	Empty() bool
	ReadRecord() (*entity.Record, error)
	Close() error
}

type FileStorageWriter interface {
	WriteRecord(record *entity.Record) error
	Close() error
}

type FileStorageEnclosure interface {
	Save(fileName string, reader io.Reader) error
	Open(fileName string) (io.Reader, error)
	Remove(fileName string) error
}

type Archive interface {
	ReadZIPByChunks(w io.Writer) error
	WriteZIPByChunks(r io.Reader) error
	Pack() error
	Unpack() error
	GetZIPModTime() (time.Time, error)
}
