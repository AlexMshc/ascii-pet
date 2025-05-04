package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"ascii-pet/internal/gen/models"
)

var ErrNotFound = errors.New("no pet data found")

type FileStore struct {
	Path string
}

func NewFileStore(path string) *FileStore {
	return &FileStore{Path: path}
}

func (fs *FileStore) Save(pet models.PetProperties) error {
	if err := os.MkdirAll(filepath.Dir(fs.Path), 0755); err != nil {
		return err
	}
	b, err := json.Marshal(pet)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.Path, b, 0644)
}

func (fs *FileStore) Load() (models.PetProperties, error) {
	var pet models.PetProperties
	b, err := os.ReadFile(fs.Path)
	if errors.Is(err, os.ErrNotExist) {
		return pet, ErrNotFound
	} else if err != nil {
		return pet, err
	}
	if err := json.Unmarshal(b, &pet); err != nil {
		return pet, err
	}
	return pet, nil
}

func (fs *FileStore) Delete() error {
	return os.Remove(fs.Path)
}
