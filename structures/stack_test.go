package structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStackNew(t *testing.T) {
	stack := NewStack[string]()

	require.NotNil(t, stack)
}

func TestStackPush1(t *testing.T) {
	stack := NewStack("data")

	ret := stack.Pop()

	require.Equal(t, "data", ret)
}

func TestStackPush2(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("check")

	ret := stack.Pop()

	require.Equal(t, "check", ret)
}

func TestStackPush3(t *testing.T) {
	stack := NewStack("A", "B")

	b := stack.Pop()
	a := stack.Pop()

	require.Equal(t, "B", b)
	require.Equal(t, "A", a)
}

func TestStackOnlyPop(t *testing.T) {
	stack := NewStack[string]()

	require.Panics(t, func() {
		stack.Pop()
	})
}

func TestStackIsEmptyWithData(t *testing.T) {
	stack := NewStack("A")

	ret := stack.IsEmpty()

	require.False(t, ret)
}

func TestStackIsEmpty(t *testing.T) {
	stack := NewStack[string]()

	ret := stack.IsEmpty()

	require.True(t, ret)
}

func TestStackPeek(t *testing.T) {
	stack := NewStack("A")

	a := stack.Peek()
	b := stack.Peek()

	require.Equal(t, "A", a)
	require.Equal(t, "A", b)
}
