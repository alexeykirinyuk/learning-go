package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(fileName string, blob []byte) (*File, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   id,
		Name: fileName,
		Data: blob,
	}, nil
}

func (s *File) String() string {
	return fmt.Sprintf("File(%s, %v)", s.Name, s.ID)
}
