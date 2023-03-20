package programmers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemoryRepository_New(t *testing.T) {
	repo := NewMemoryRepository()
	require.NotNil(t, repo)
}

func TestMemoryRepository_IsExists(t *testing.T) {
	repo := NewMemoryRepository()

	ret := repo.IsExists("abc")

	require.Equal(t, false, ret)
}

func TestMemoryRepository_Save(t *testing.T) {
	repo := NewMemoryRepository()

	err := repo.Save("abc", []byte("abc"))

	require.NoError(t, err)
}

func TestMemoryRepository_IsExistsTwoFiles(t *testing.T) {
	repo := NewMemoryRepository()
	_ = repo.Save("abc", []byte("abc"))

	abcExists := repo.IsExists("abc")
	zxcExists := repo.IsExists("zxc")

	require.Equal(t, true, abcExists)
	require.Equal(t, false, zxcExists)
}

func TestMemoryRepository_Load(t *testing.T) {
	repo := NewMemoryRepository()
	_ = repo.Save("abc", []byte("abc"))

	bytes, err := repo.Load("abc")

	require.NoError(t, err)
	require.Equal(t, []byte("abc"), bytes)
}

func TestMemoryRepository_Load_(t *testing.T) {
	repo := NewMemoryRepository()

	bytes, err := repo.Load("abc")

	require.Error(t, err)
	require.Nil(t, bytes)
}

func TestMemoryRepository_List(t *testing.T) {
	repo := NewMemoryRepository()
	_ = repo.Save("abc", []byte("abc"))

	files, err := repo.List()

	require.NoError(t, err)
	require.Equal(t, 1, len(files))
	require.Equal(t, "abc", files[0])
}
