package repo

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/repo/file"
	"sync"
)

// Repo store in - memory storage and file - writer.
type Repo struct {
	mu     sync.Mutex
	record *entity.Record
	writer *file.Writer
}

// NewFileRepo create repository to store information in file.
func NewFileRepo(filePath string) (*Repo, error) {
	records, err := createRecords(filePath)
	if err != nil {
		return nil, err
	}
	writer, err := file.NewWriter(filePath)
	if err != nil {
		return nil, err
	}

	return &Repo{
		record: records,
		writer: writer,
	}, nil
}

// Close Closes file writer if information stored in file.
func (r *Repo) Close() error {
	return r.writer.Close()
}

func createRecords(filePath string) (*entity.Record, error) {
	reader, err := file.NewReader(filePath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	record, err := reader.ReadRecord()
	if err != nil {
		return nil, err
	}
	return record, nil
}
