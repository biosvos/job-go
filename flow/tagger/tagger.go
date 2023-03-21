package tagger

type Tagger interface {
	Tagging(content string) []string
}
