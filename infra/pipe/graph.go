package pipe

type Accessor string

type Graph struct {
	elements map[string]*vertex
}

type vertex struct {
	linker map[Accessor][]*vertex
	value  string
}

func newVertex(value string) *vertex {
	return &vertex{
		linker: make(map[Accessor][]*vertex),
		value:  value,
	}
}

func NewGraph() *Graph {
	return &Graph{
		elements: make(map[string]*vertex),
	}
}

func (g *Graph) AddVertex(node string) {
	g.elements[node] = newVertex(node)
}

func (g *Graph) Link(from string, to string, accessor Accessor) {
	fromVertex, ok := g.elements[from]
	if !ok {
		return
	}

	toVertex, ok := g.elements[to]
	if !ok {
		return
	}

	fromVertex.linker[accessor] = append(fromVertex.linker[accessor], toVertex)
}

func (g *Graph) HasVertex(node string) bool {
	_, ok := g.elements[node]
	return ok
}

func (g *Graph) ListLinked(node string, accessor Accessor) []string {
	fromVertex, ok := g.elements[node]
	if !ok {
		return nil
	}
	vertices := fromVertex.linker[accessor]
	var ret []string
	for _, to := range vertices {
		ret = append(ret, to.value)
	}
	return ret
}

func (g *Graph) List() []string {
	var ret []string
	for key := range g.elements {
		ret = append(ret, key)
	}
	return ret
}
