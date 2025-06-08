package rules

import (
  "regexp"
  "sandhi_transformer/sandhi/core"
)

var consonantRules = []core.SandhiRule{
  {
    Pattern:     regexp.MustCompile(`(t) (c)`),
    Explanation: "t + c → cc (tat + cit → taccit)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[cc]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(t) (j)`),
    Explanation: "t + j → jj (tat + jana → tajjana)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[jj]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(d) (dh)`),
    Explanation: "d + dh → ddh (sad + dharma → saddharma)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ddh]", 1, 2) },
  },
  {
    Pattern:     regexp.MustCompile(`(n) (c)`),
    Explanation: "n + c → ñc (rājan + candra → rājañcandra)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ñc]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(n) (j)`),
    Explanation: "n + j → ñj (rājan + jñāna → rājañjñāna)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ñj]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(n) (ṭ)`),
    Explanation: "n + ṭ → ṇṭ (rājan + ṭīkā → rājaṇṭīkā)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[ṇṭ]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(n) (t)`),
    Explanation: "n + t → nt (rājan + tapaḥ → rājantapaḥ)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[nt]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(n) (p)`),
    Explanation: "n + p → mp (rājan + patyate → rājampatyate)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[mp]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(t) (ś)`),
    Explanation: "t + ś → cch (tat + śivaḥ → tacchivaḥ)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[cch]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(s) (ca)`),
    Explanation: "s + ca → śca (rāmas + ca → rāmaśca)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[śca]", 1, 2) },
  },
  {
    Pattern:     regexp.MustCompile(`(t) (s)`),
    Explanation: "t + s → s (tat + satyam → tas satyam)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[s] s", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(s) (k)`),
    Explanation: "s + k → sk",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[sk]", 1, 1) },
  },
  {
    Pattern:     regexp.MustCompile(`(s) ([aāiīuūṛṝḷḹeoēōghjbdgz])`),
    Explanation: "s → r before voiced letters (manas + indra → manarindraḥ)",
    ReplaceFunc: func(_, g1, g2 string) string { return core.JoinSandhi(g1, g2, "[r]", 1, 0) },
  },
}

type ConsonantTransformer struct{}

func (t ConsonantTransformer) TransformToChunks(input string) []core.Chunk {
  return core.ApplySandhiRules(input, consonantRules)
}
