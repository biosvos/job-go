package pipe

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGraph_New(t *testing.T) {
	graph := NewGraph()
	require.NotNil(t, graph)
}

func TestGraph_HasVertex(t *testing.T) {
	graph := NewGraph()

	ret := graph.HasVertex("abc")

	require.Equal(t, false, ret)
}

func TestGraph_Link(t *testing.T) {
	graph := NewGraph()

	graph.Link("from", "to", "access")
}

func TestGraph_AddHasVertex(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("abc")

	ret := graph.HasVertex("abc")

	require.Equal(t, true, ret)
}

func TestGraph_ListLinked(t *testing.T) {
	graph := NewGraph()

	ret := graph.ListLinked("abc", "accessor")

	require.Nil(t, ret)
}

func TestGraph_AddAddListLinked(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.Link("A", "B", "link")

	strings := graph.ListLinked("A", "link")

	require.Equal(t, 1, len(strings))
	require.Equal(t, "B", strings[0])
}

func TestGraph_List(t *testing.T) {
	graph := NewGraph()

	strings := graph.List()

	require.Equal(t, 0, len(strings))
}

func TestGraph_AddList(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("ABC")

	strings := graph.List()

	require.Equal(t, 1, len(strings))
	require.Equal(t, "ABC", strings[0])
}
