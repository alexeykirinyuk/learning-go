package storage

import (
	"github.com/alexeykirinyuk/learning-go/2-workshop/storage/internal/file"
)

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error) {
	return file.NewFile(fileName, blob)
	// if err != nil {
	// 	return nil, err
	// }

	// return file, nil
}
