package pipe

import "job-go/flow/tagger"

var _ tagger.Tagger = &Pipe{}

type Pipe struct{}

func NewPipeTagger(path string) *Pipe {
	return &Pipe{}
}

func (t *Pipe) Tagging(content string) []*tagger.Tag {
	if content == "" {
		return nil
	}
	return []*tagger.Tag{
		{
			Names: []string{content},
		},
	}
}
