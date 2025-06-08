package core

import "regexp"

type SandhiRule struct {
  Pattern     *regexp.Regexp
  Explanation string
  ReplaceFunc func(fullMatch, group1, group2 string) string
}

type Chunk struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip,omitempty"`
}

type ChunkingTransformer interface {
	TransformToChunks(input string) []Chunk
}
