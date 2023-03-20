package pipe

import (
	"github.com/stretchr/testify/require"
	"job-go/flow/tagger"
	"testing"
)

const testTagFile = "file.tag"

func TestPipe_New(t *testing.T) {
	pipe := NewPipeTagger(testTagFile)

	require.NotNil(t, pipe)
}

func TestPipe_EmptyTag(t *testing.T) {
	pipe := NewPipeTagger(testTagFile)

	tags := pipe.Tagging("")

	require.Equal(t, 0, len(tags))
}

func TestPipe_GoTag(t *testing.T) {
	pipe := NewPipeTagger(testTagFile)

	tags := pipe.Tagging("Go")

	require.Equal(t, 1, len(tags))
	require.Equal(t, &tagger.Tag{Names: []string{"Go"}}, tags[0])
}

func TestPipe_AliasGoTag(t *testing.T) {
	pipe := NewPipeTagger(testTagFile)

	tags := pipe.Tagging("go")

	require.Equal(t, 1, len(tags))
	require.Equal(t, &tagger.Tag{Names: []string{"Go"}}, tags[0])
}
