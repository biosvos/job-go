package programmers

import (
	"github.com/pkg/errors"
	"os"
)

var _ Repository = &FileRepository{}

type FileRepository struct {
	prefix string
}

func NewFileRepository(prefix string) *FileRepository {
	return &FileRepository{
		prefix: prefix,
	}
}

func (f *FileRepository) IsExists(path string) bool {
	path = f.prefix + "/" + path
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func (f *FileRepository) Save(path string, bytes []byte) error {
	path = f.prefix + "/" + path
	err := os.WriteFile(path, bytes, 0600)
	return errors.WithStack(err)
}

func (f *FileRepository) Load(path string) ([]byte, error) {
	path = f.prefix + "/" + path
	ret, err := os.ReadFile(path)
	return ret, err
}
