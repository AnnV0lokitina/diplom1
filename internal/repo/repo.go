package repo

import (
	"sync"
)

type Repo struct {
	mu        sync.Mutex
	storePath string
}
