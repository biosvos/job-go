package programmers

type Repository interface {
	IsExists(path string) bool
	Save(path string, bytes []byte) error
	Load(path string) ([]byte, error)
}
