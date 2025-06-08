package rules

import (
	"regexp"
  "sandhi_transformer/sandhi/core"
)

var visargaRules = []core.SandhiRule{
  {
    Pattern:     regexp.MustCompile(`(ḥ) (p)`), // also "ph"
    Explanation: "‘ḥ’ becomes ‘f’ (upadhmānīya) before voiceless bilabial stops (p/ph)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[f] ", 1, 0) },
  },
  {
    Pattern:     regexp.MustCompile(`(ḥ) (k)`), // also "kh"
    Explanation: "‘ḥ’ becomes ‘H’ (jihvāmūlīya) before voiceless velar stops (k/kh)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[H] ", 1, 0) },
  },
  {
    Pattern:     regexp.MustCompile(`(ḥ) ([śṣs])`),
    Explanation: "‘ḥ’ is dropped before sibilants (ś/ṣ/s)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[] ", 1, 0) },
  },
  // {
  //   Pattern:     regexp.MustCompile(`(ḥ) ([aāiīuūṛṝḷḹeoēō])`),
  //   Explanation: "‘ḥ’ becomes ‘ō’ before a vowel",
  //   ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ō] ", 2, 1) },
  // },
  // {
  //   Pattern:     regexp.MustCompile(`([aāiīuūṛṝḷḹeoēō])ḥ([^\w\n\r]*)`),
  //   Explanation: "Word-final visarga is retained with marker [ḥV] to indicate echoing the vowel.",
  //   ReplaceFunc: func(_, g1, g2 string) string {
  // 	if len(g1) < 2 {
  // 		return g1 + g2
  // 	}
  // 	runes := []rune(g1)
  // 	vowel := string(runes[len(runes)-2])
  // 	return g1[:len(g1)-1] + "[ḥ" + vowel + "]" + g2
  // },
  // },
}

type VisargaTransformer struct{}

func (t VisargaTransformer) TransformToChunks(input string) ([]core.Chunk) {
  return core.ApplySandhiRules(input, visargaRules)
}

// https://learnsanskrit.org/guide/core/sandhi/
// Another change is that the visarga disappears if a voiced sound follows it:
// But one important exception is that aḥ becomes o if a voiced consonant follows:
