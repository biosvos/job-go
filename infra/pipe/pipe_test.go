package pipe

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testTagFile = "file.tag"

func TestPipe_New(t *testing.T) {
	pipe, err := NewPipeTagger(testTagFile)

	require.NoError(t, err)
	require.NotNil(t, pipe)
}

func TestPipe_EmptyTag(t *testing.T) {
	pipe, _ := NewPipeTagger(testTagFile)

	tags := pipe.Tagging("")

	require.Equal(t, 0, len(tags))
}

func TestPipe_GoTag(t *testing.T) {
	pipe, _ := NewPipeTagger(testTagFile)

	tags := pipe.Tagging("Go")

	require.Equal(t, 1, len(tags))
	require.Equal(t, "Go", tags[0])
}

func TestPipe_AliasGoTag(t *testing.T) {
	pipe, _ := NewPipeTagger(testTagFile)

	tags := pipe.Tagging("go")

	require.Equal(t, 1, len(tags))
	require.Equal(t, "Go", tags[0])
}
