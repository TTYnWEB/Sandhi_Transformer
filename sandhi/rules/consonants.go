package rules

import (
	"regexp"
  "sandhi_transformer/sandhi/core"
)

var consonantRules = []core.SandhiRule{
  {
    Pattern:     regexp.MustCompile(`t\s+c`),
    Explanation: "t + c → cc (tat + cit → taccit)",
    ReplaceFunc: func(sub []string) string { return "[cc]" },
  },
  {
    Pattern:     regexp.MustCompile(`t\s+j`),
    Explanation: "t + j → jj (tat + jana → tajjana)",
    ReplaceFunc: func(sub []string) string { return "[jj]" },
  },
  {
    Pattern:     regexp.MustCompile(`d\s+dh`),
    Explanation: "d + dh → ddh (sad + dharma → saddharma)",
    ReplaceFunc: func(sub []string) string { return "[ddh]" },
  },
  {
    Pattern:     regexp.MustCompile(`n\s+c`),
    Explanation: "n + c → ñc (rājan + candra → rājañcandra)",
    ReplaceFunc: func(sub []string) string { return "[ñc]" },
  },
  {
    Pattern:     regexp.MustCompile(`n\s+j`),
    Explanation: "n + j → ñj (rājan + jñāna → rājañjñāna)",
    ReplaceFunc: func(sub []string) string { return "[ñj]" },
  },
  {
    Pattern:     regexp.MustCompile(`n\s+ṭ`),
    Explanation: "n + ṭ → ṇṭ (rājan + ṭīkā → rājaṇṭīkā)",
    ReplaceFunc: func(sub []string) string { return "[ṇṭ]" },
  },
  {
    Pattern:     regexp.MustCompile(`n\s+t`),
    Explanation: "n + t → nt (rājan + tapaḥ → rājantapaḥ)",
    ReplaceFunc: func(sub []string) string { return "[nt]" },
  },
  {
    Pattern:     regexp.MustCompile(`n\s+p`),
    Explanation: "n + p → mp (rājan + patyate → rājampatyate)",
    ReplaceFunc: func(sub []string) string { return "[mp]" },
  },
  {
    Pattern:     regexp.MustCompile(`t\s+ś`),
    Explanation: "t + ś → cch (tat + śivaḥ → tacchivaḥ)",
    ReplaceFunc: func(sub []string) string { return "[cch]" },
  },
  {
    Pattern:     regexp.MustCompile(`s\s+ca`),
    Explanation: "s + ca → śca (rāmas + ca → rāmaśca)",
    ReplaceFunc: func(sub []string) string { return "[śca]" },
  },
  {
    Pattern:     regexp.MustCompile(`t\s+s`),
    Explanation: "t + s → s (tat + satyam → tas satyam)",
    ReplaceFunc: func(sub []string) string { return "[s] s" },
  },
	{
    Pattern:     regexp.MustCompile(`s\s+k`),
    Explanation: "s + k → sk",
    ReplaceFunc: func(sub []string) string { return "[sk]" },
	},
  {
    Pattern:     regexp.MustCompile(`s\s+([aāiīuūṛṝḷḹeoēōghjbdgz])`),
    Explanation: "s → r before voiced letters (manas + indra → manarindraḥ)",
    ReplaceFunc: func(sub []string) string { return "[r]" + sub[1] },
  },
}

type ConsonantTransformer struct{}

func (t ConsonantTransformer) Transform(input string) string {
	out, final := t.TransformWithExplanation(input)
  if (len(out) == 0) {
    return input
  } else {
		return final
	}
}

func (t ConsonantTransformer) TransformWithExplanation(input string) ([]core.Transformation, string) {
  return core.ApplyRules(input, consonantRules)
}
