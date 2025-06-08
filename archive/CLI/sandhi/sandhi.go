package sandhi

import (
  "sandhi_transformer/sandhi/core"
  "sandhi_transformer/sandhi/rules"
)

var transformers = []core.Transformer{
  rules.AnusvaraTransformer{},
  rules.ConsonantTransformer{},
  rules.VisargaTransformer{},
	rules.VowelTransformer{},
}

var explainingTransformers = []core.ExplainingTransformer{
  rules.AnusvaraTransformer{},
  rules.ConsonantTransformer{},
  rules.VisargaTransformer{},
	rules.VowelTransformer{},
}

// Applies all registered Transformers (no explanations)
func ApplyAllSandhi(input string) string {
  for _, t := range transformers {
    input = t.Transform(input)
  }
  return input
}

// Applies all ExplainingTransformers and gathers explanations
func ApplyAllSandhiWithExplanation(input string) ([]core.Transformation, string) {
  all := []core.Transformation{}
  for _, t := range explainingTransformers {
    tr, out := t.TransformWithExplanation(input)
    all = append(all, tr...)
    input = out
  }
  return all, input
}
