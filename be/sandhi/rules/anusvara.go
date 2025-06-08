package rules

import (
	"regexp"
  "sandhi_transformer/sandhi/core"
)
// nasalMap maps consonants to their corresponding nasal.
var nasalMap = map[string]string{
  "kh": "ṅ",
  "k":  "ṅ",
  "gh": "ṅ",
  "g":  "ṅ",
  "ṅ":  "ṅ",
  "ch": "ñ",
  "c":  "ñ",
  "jh": "ñ",
  "j":  "ñ",
  "ñ":  "ñ",
  "ṭh": "ṇ",
  "ṭ":  "ṇ",
  "ḍh": "ṇ",
  "ḍ":  "ṇ",
  "ṇ":  "ṇ",
  "th": "n",
  "t":  "n",
  "dh": "n",
  "d":  "n",
  "n":  "n",
  "ph": "m",
  "p":  "m",
  "bh": "m",
  "b":  "m",
  "m":  "m",
}

func replaceFunc(_, g1, g2 string) string {
  first := ""

	if len(g2) >= 2 {
		if _, ok := nasalMap[g2[:2]]; ok {
			first = g2[:2]
		} else {
			first = string([]rune(g2)[0])
		}
	} else if len(g2) == 1 {
		first = g2
	}

	nasal, ok := nasalMap[first]
	if !ok {
		nasal = "ṁ" // fallback
	}

	sub := "[" + nasal + "] "
	return core.JoinSandhi(g1, g2, sub, 1, 0)
}

var anusvaraRules = []core.SandhiRule{
  {
    Pattern:     regexp.MustCompile(`(ṁ) ([ckgjṭḍtdpbṅñṇnm]h?)`),
    Explanation: "Anusvāra [ṁ] changes to the homorganic nasal before a consonant.",
    ReplaceFunc:  replaceFunc,
  },
}

type AnusvaraTransformer struct{}

func (t AnusvaraTransformer) TransformToChunks(input string) []core.Chunk {
	return core.ApplySandhiRules(input, anusvaraRules)
}
