package pipe

import (
	"job-go/flow/tagger"
	"job-go/structures"
	"os"
	"sort"
	"strings"
)

var _ tagger.Tagger = &Pipe{}

type Pipe struct {
	graph *Graph
}

const (
	representativeAccessor = Accessor("representative")
	parentsAccessor        = Accessor("parents")
	childrenAccessor       = Accessor("children")
)

func NewPipeTagger(path string) (*Pipe, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	graph := NewGraph()
	stack := structures.NewStack[string]()
	lines := strings.Split(strings.ReplaceAll(string(bytes), "\r\n", "\n"), "\n")
	for _, line := range lines {
		count := strings.Count(line, "\t")
		line = strings.ReplaceAll(line, "\t", "")
		words := strings.Split(line, "|")
		for stack.Size() < count {
			stack.Pop()
		}
		representative := words[0]
		for _, word := range words {
			graph.AddVertex(word)
			if !stack.IsEmpty() {
				graph.Link(stack.Peek(), word, childrenAccessor)
				graph.Link(word, stack.Peek(), parentsAccessor)
			}
			graph.Link(word, representative, representativeAccessor)
		}
		stack.Push(representative)
	}

	return &Pipe{
		graph: graph,
	}, nil
}

func (p *Pipe) Tagging(content string) []string {
	if content == "" {
		return nil
	}
	tags := p.graph.List()
	sort.Slice(tags, func(i, j int) bool {
		return len(tags[i]) > len(tags[j])
	})

	var ret []string
	for _, tag := range tags {
		if strings.Contains(content, tag) {
			content = strings.ReplaceAll(content, tag, "")
			linked := p.graph.ListLinked(tag, representativeAccessor)
			ret = append(ret, linked[0])
		}
	}
	return ret
}
