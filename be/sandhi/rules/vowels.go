package rules

import (
	"regexp"
  "sandhi_transformer/sandhi/core"
)

var vowelRules = []core.SandhiRule{
  {
    Pattern:     regexp.MustCompile(`(a) (a)`),
    Explanation: "a + a → ā (savarna dīrga)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ā]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(i) (i)`),
    Explanation: "i + i → ī (savarna dīrga)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ī]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(u) (u)`),
    Explanation: "u + u → ū (savarna dīrga)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ū]", 1, 1) },
  },
	{
    Pattern:     regexp.MustCompile(`([aā]) ([iī])`),
    Explanation: "a/ā + i/ī → ē (guṇa)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ē]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`([aā]) ([uū])`),
    Explanation: "a/ā + u/ū → ō (guṇa)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ō]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`([aā]) (ē)`),
    Explanation: "a/ā + ē → ai (vṛddhi)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ai]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`([aā]) (ō)`),
    Explanation: "a/ā + ō → au (vṛddhi)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[au]", 1, 1) },
  },
}

type VowelTransformer struct{}

func (t VowelTransformer) TransformToChunks(input string) []core.Chunk {
	return core.ApplySandhiRules(input, vowelRules)
}

/*
{
  Pattern:     regexp.MustCompile(`([aāēō])\s+(i|ī)`),
  Explanation: "(yaṇ)",
  // ReplaceFunc: func(sub []string) string { return sub[1] + "y" + sub[2] },
  // ReplaceFunc: func(sub []string) string { return "[au]" },
},
{
  Pattern:     regexp.MustCompile(`([aāēō])\s+(u|ū)`),
  Explanation: "(yaṇ)",
  // ReplaceFunc: func(sub []string) string { return sub[1] + "v" + sub[2] },
},
{
  Pattern:     regexp.MustCompile(`([aāēō])\s+ṛ`),
  Explanation: "(yaṇ)",
  // ReplaceFunc: func(sub []string) string { return sub[1] + "r" + sub[2] },
},
*/
