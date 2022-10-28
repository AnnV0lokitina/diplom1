package repo

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"sync"
)

// Repo store in - memory storage and file - writer.
type Repo struct {
	mu        sync.Mutex
	record    *entity.Record
	writer    FileStorageWriter
	reader    FileStorageReader
	enclosure FileStorageEnclosure
	archive   Archive
}

// NewFileRepo create repository to store information in file.
func NewFileRepo(
	reader FileStorageReader,
	writer FileStorageWriter,
	enclosure FileStorageEnclosure,
	archive Archive,
) (*Repo, error) {
	var (
		record *entity.Record
		err    error
	)
	if !reader.Empty() {
		defer func() {
			err = reader.Close()
		}()
		record, err = reader.ReadRecord()
		if err != nil {
			return nil, err
		}
	} else {
		record = &entity.Record{}
	}
	return &Repo{
		record:    record,
		writer:    writer,
		enclosure: enclosure,
		archive:   archive,
		reader:    reader,
	}, err
}

// Close Closes file writer if information stored in file.
func (r *Repo) Close() error {
	return r.writer.Close()
}
