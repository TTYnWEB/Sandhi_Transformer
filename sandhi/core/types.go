package core

import "regexp"

type SandhiRule struct {
  Pattern     *regexp.Regexp
  Explanation string
  ReplaceFunc func([]string) string
}

type Transformer interface {
  Transform(input string) string
}

type ExplainingTransformer interface {
  TransformWithExplanation(input string) ([]Transformation, string)
}

type Transformation struct {
  Original    string
  Transformed string
  Explanation string
}
