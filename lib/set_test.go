package lib

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetNew(t *testing.T) {
	set := NewSet[int]()
	require.NotNil(t, set)
}

func TestSetAdd(t *testing.T) {
	set := NewSet[int]()
	set.Add(3)

	ret := set.Exists(3)

	require.True(t, ret)
}

func TestSetExists(t *testing.T) {
	set := NewSet[int]()

	ret := set.Exists(3)

	require.False(t, ret)
}

func TestSetDelete(t *testing.T) {
	set := NewSet[int]()
	set.Add(3)
	set.Delete(3)

	exists := set.Exists(3)

	require.False(t, exists)
}

func TestSetList(t *testing.T) {
	set := NewSet[int]()
	set.Add(3)
	set.Add(4)

	slice := set.Slice()

	require.Equal(t, 2, len(slice))
}
