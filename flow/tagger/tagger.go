package tagger

type Tag struct {
	Names []string
}

type Tagger interface {
	Tagging(content string) []*Tag
}
