package programmers

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func createFileRepository(t *testing.T) (*FileRepository, func()) {
	temp, err := os.MkdirTemp("", "**********")
	if err != nil {
		t.Fatal(err)
	}
	return NewFileRepository(temp), func() {
		_ = os.RemoveAll(temp)
	}
}

func TestFileRepository_New(t *testing.T) {
	repo, closer := createFileRepository(t)
	defer closer()

	require.NotNil(t, repo)
}

func TestFileRepository_IsExists(t *testing.T) {
	repo, closer := createFileRepository(t)
	defer closer()

	ret := repo.IsExists("abc")

	require.Equal(t, false, ret)
}

func TestFileRepository_Save(t *testing.T) {
	repo, closer := createFileRepository(t)
	defer closer()

	err := repo.Save("abc", []byte("abc"))

	require.NoError(t, err)
}

func TestFileRepository_IsExistsTwoFiles(t *testing.T) {
	repo, closer := createFileRepository(t)
	defer closer()
	_ = repo.Save("abc", []byte("abc"))

	abcExists := repo.IsExists("abc")
	zxcExists := repo.IsExists("zxc")

	require.Equal(t, true, abcExists)
	require.Equal(t, false, zxcExists)
}

func TestFileRepository_Load(t *testing.T) {
	repo, closer := createFileRepository(t)
	defer closer()
	_ = repo.Save("abc", []byte("abc"))

	bytes, err := repo.Load("abc")

	require.NoError(t, err)
	require.Equal(t, []byte("abc"), bytes)
}

func TestFileRepository_Load_(t *testing.T) {
	repo, closer := createFileRepository(t)
	defer closer()

	bytes, err := repo.Load("abc")

	require.Error(t, err)
	require.Nil(t, bytes)
}
