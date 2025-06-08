package sandhi

import (
  "sandhi_transformer/sandhi/core"
  "sandhi_transformer/sandhi/rules"
)

var explainingTransformers = []core.ChunkingTransformer{
  rules.AnusvaraTransformer{},
  rules.ConsonantTransformer{},
  rules.VisargaTransformer{},
  rules.VowelTransformer{},
}

func ApplyAllSandhiChunks(input string) []core.Chunk {
  current := input
  var finalChunks []core.Chunk

  for _, transformer := range explainingTransformers {
    chunks := transformer.TransformToChunks(current)

    // Check if any chunk has a tooltip → if so, it's a match
    hasChange := false
    for _, ch := range chunks {
      if ch.Tooltip != "" {
        hasChange = true
        break
      }
    }

    if hasChange {
      // Overwrite finalChunks with current transformation result
      finalChunks = chunks

      // Concatenate all chunk.Text for the next round
      next := ""
      for _, ch := range chunks {
        next += ch.Text
      }
      current = next
    }
  }

  // If no rules matched at all
  if len(finalChunks) == 0 {
    return []core.Chunk{{Text: input}}
  }

  return finalChunks
}
