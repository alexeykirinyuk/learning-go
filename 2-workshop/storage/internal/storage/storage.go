package storage

import (
	"fmt"

	"github.com/alexeykirinyuk/learning-go/2-workshop/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	fileMap map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		fileMap: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error) {
	file, err := file.NewFile(fileName, blob)
	if err != nil {
		return nil, err
	}

	s.fileMap[file.ID] = file

	return file, nil
}

func (s *Storage) GetByID(id uuid.UUID) (*file.File, error) {
	item, ok := s.fileMap[id]
	if !ok {
		return nil, fmt.Errorf("file '%v' not found", id)
	}

	return item, nil
}