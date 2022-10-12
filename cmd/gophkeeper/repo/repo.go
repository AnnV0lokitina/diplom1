package repo

import (
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"os"
	"path/filepath"
	"sync"
)

const dataFileName = "data.json"

// Repo store in - memory storage and file - writer.
type Repo struct {
	mu        sync.Mutex
	record    *entity.Record
	writer    *entity.Writer
	storePath string
}

// NewFileRepo create repository to store information in file.
func NewFileRepo(dir string) (*Repo, error) {
	filePath := filepath.Join(dir, dataFileName)
	records, err := createRecords(filePath)
	if err != nil {
		return nil, err
	}
	writer, err := entity.NewWriter(filePath)
	if err != nil {
		return nil, err
	}

	return &Repo{
		record:    records,
		writer:    writer,
		storePath: dir,
	}, nil
}

// Close Closes file writer if information stored in file.
func (r *Repo) Close() error {
	return r.writer.Close()
}

func createRecords(filePath string) (*entity.Record, error) {
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) || stat.Size() == 0 {
		return &entity.Record{}, nil
	}
	reader, err := entity.NewReader(filePath)
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
