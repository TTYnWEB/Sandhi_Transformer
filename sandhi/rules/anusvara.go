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

var anusvaraRules = []core.SandhiRule{
  {
    Pattern:     regexp.MustCompile(`ṁ(\s*)([ckgjṭḍtdpbṅñṇnm])(h?)`),
    Explanation: "Anusvāra [ṁ] changes to the homorganic nasal before a consonant.",
    ReplaceFunc: func(sub []string) string {
      space := sub[1]
      cons := sub[2]
      asp := sub[3]
      fullCons := cons + asp

      nasal, ok := nasalMap[fullCons]
      if !ok {
        return "ṁ" + space + fullCons
      }
      return "[" + nasal + "]" + space + fullCons
    },
  },
}

type AnusvaraTransformer struct{}

func (t AnusvaraTransformer) Transform(input string) string {
	out, final := t.TransformWithExplanation(input)
  if (len(out) == 0) {
    return input
  } else {
		return final
	}
}

func (t AnusvaraTransformer) TransformWithExplanation(input string) ([]core.Transformation, string) {
  return core.ApplyRules(input, anusvaraRules)
}

// func (t AnusvaraTransformer) TransformWithExplanation(input string) ([]sandhi.Transformation, string) {
//   transformations := []sandhi.Transformation{}
//   result := input
//
//   result = anusvaraRe.ReplaceAllStringFunc(result, func(match string) string {
//     sub := anusvaraRe.FindStringSubmatch(match)
//     if len(sub) < 4 {
//       return match
//     }
//     space := sub[1]
//     cons := sub[2]
//     asp := sub[3]
//     fullCons := cons + asp
//
//     nasal, ok := nasalMap[fullCons]
//     if !ok {
//       return match
//     }
//
//     transformed := "[" + nasal + "]" + space + fullCons
// 		explanation := "Anusvāra [ṁ ] changes to the homorganic nasal [" + nasal + "] before the consonant [" + fullCons + "]."
//     transformations = append(transformations, sandhi.Transformation{
//       Original:    match,
//       Transformed: transformed,
//       Explanation: explanation,
//     })
//
//     return transformed
//   })
//
//   return transformations, result
// }
//
// // func TransformAnusvara(text string) string {
// //     return anusvaraRe.ReplaceAllStringFunc(text, func(match string) string {
// //         sub := anusvaraRe.FindStringSubmatch(match)
// //         if len(sub) < 4 {
// //             return match
// //         }
// //         space := sub[1]
// //         cons := sub[2]
// //         asp := sub[3]
// //         key := cons + asp
// //         if nasal, ok := nasalMap[key]; ok {
// //             return "[" + nasal + "]" + space + key
// //         }
// //         return match
// //     })
// // }
