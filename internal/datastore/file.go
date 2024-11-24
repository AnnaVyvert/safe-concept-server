package datastore

import (
	"fmt"
	"os"
	"path"
)

var _ Storage = fileStorage{}

type fileStorage struct {
	storageDir string
}

func NewFileStorage(storageDir string) Storage {
	if err := os.MkdirAll(storageDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("can not make dir %q: %q", storageDir, err))
	}
	return fileStorage{storageDir: storageDir}
}

// Load implements Storage.
func (f fileStorage) Load(id Identifier) ([]byte, error) {
	filePath := path.Join(f.storageDir, id.Identity())
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Store implements Storage.
func (f fileStorage) Store(id Identifier, data []byte) error {
	filePath := path.Join(f.storageDir, id.Identity())
	return os.WriteFile(filePath, data, os.ModePerm)
}
