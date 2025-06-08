package rules

import (
	"regexp"
  "sandhi_transformer/sandhi/core"
)

var vowelRules = []core.SandhiRule{
  {
    Pattern:     regexp.MustCompile(`a\s+a`),
    Explanation: "a + a → ā (savarna dīrga)",
    ReplaceFunc: func(sub []string) string { return "[ā]" },
  },
  {
    Pattern:     regexp.MustCompile(`i\s+i`),
    Explanation: "i + i → ī (savarna dīrga)",
    ReplaceFunc: func(sub []string) string { return "[ī]" },
  },
  {
    Pattern:     regexp.MustCompile(`u\s+u`),
    Explanation: "u + u → ū (savarna dīrga)",
    ReplaceFunc: func(sub []string) string { return "[ū]" },
  },
	{
    Pattern:     regexp.MustCompile(`[aā]\s+[iī]`),
    Explanation: "a/ā + i/ī → ē (guṇa)",
    ReplaceFunc: func(sub []string) string { return "[ē]" },
  },
  {
    Pattern:     regexp.MustCompile(`[aā]\s+[uū]`),
    Explanation: "a/ā + u/ū → ō (guṇa)",
    ReplaceFunc: func(sub []string) string { return "[ō]" },
  },
  {
    Pattern:     regexp.MustCompile(`[aā]\s+ē`),
    Explanation: "a/ā + ē → ai (vṛddhi)",
    ReplaceFunc: func(sub []string) string { return "[ai]" },
  },
  {
    Pattern:     regexp.MustCompile(`[aā]\s+ō`),
    Explanation: "a/ā + ō → au (vṛddhi)",
    ReplaceFunc: func(sub []string) string { return "[au]" },
  },
}

type VowelTransformer struct{}

func (t VowelTransformer) Transform(input string) string {
	out, final := t.TransformWithExplanation(input)
  if (len(out) == 0) {
    return input
  } else {
		return final
	}
}

func (t VowelTransformer) TransformWithExplanation(input string) ([]core.Transformation, string) {
  return core.ApplyRules(input, vowelRules)
}

  // {
  //   Pattern:     regexp.MustCompile(`([aāēō])\s+(i|ī)`),
  //   Explanation: "(yaṇ)",
  // // ReplaceFunc: func(sub []string) string { return sub[1] + "y" + sub[2] },
  //   // ReplaceFunc: func(sub []string) string { return "[au]" },
  // },
  // {
  //   Pattern:     regexp.MustCompile(`([aāēō])\s+(u|ū)`),
  //   Explanation: "(yaṇ)",
  // // ReplaceFunc: func(sub []string) string { return sub[1] + "v" + sub[2] },
  // },
  // {
  //   Pattern:     regexp.MustCompile(`([aāēō])\s+ṛ`),
  //   Explanation: "(yaṇ)",
  // // ReplaceFunc: func(sub []string) string { return sub[1] + "r" + sub[2] },
  // },
