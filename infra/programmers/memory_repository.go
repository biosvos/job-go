package programmers

import "github.com/pkg/errors"

type MemoryRepository struct {
	files map[string][]byte
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		files: make(map[string][]byte),
	}
}

func (r *MemoryRepository) IsExists(path string) bool {
	_, ok := r.files[path]
	return ok
}

func (r *MemoryRepository) Save(path string, bytes []byte) error {
	r.files[path] = bytes
	return nil
}

func (r *MemoryRepository) Load(path string) ([]byte, error) {
	bytes, ok := r.files[path]
	if !ok {
		return nil, errors.New("failed to load")
	}
	return bytes, nil
}

func (r *MemoryRepository) List() ([]string, error) {
	var ret []string
	for key := range r.files {
		ret = append(ret, key)
	}
	return ret, nil
}
