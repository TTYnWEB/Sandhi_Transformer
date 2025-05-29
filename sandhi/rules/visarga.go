package rules

import (
	"regexp"
  "sandhi_transformer/sandhi/core"
)

var visargaRules = []core.SandhiRule{
  {
    Pattern:     regexp.MustCompile(`ḥ\s+(p|ph)`),
    Explanation: "‘ḥ’ becomes ‘f’ (upadhmānīya) before voiceless bilabial stops (p/ph)",
    ReplaceFunc: func(sub []string) string { return "[f] " + sub[1] },
  },
  {
    Pattern:     regexp.MustCompile(`ḥ\s+(k|kh)`),
    Explanation: "‘ḥ’ becomes ‘H’ (jihvāmūlīya) before voiceless velar stops (k/kh)",
    ReplaceFunc: func(sub []string) string { return "[H] " + sub[1] },
  },
  {
    Pattern:     regexp.MustCompile(`ḥ\s+([aāiīuūṛṝḷḹeoēō])`),
    Explanation: "‘ḥ’ becomes ‘ō’ before a vowel",
    ReplaceFunc: func(sub []string) string { return "[ō] " + sub[1] },
  },
  {
    Pattern:     regexp.MustCompile(`ḥ\s+([śṣs])`),
    Explanation: "‘ḥ’ is dropped before sibilants (ś/ṣ/s)",
    ReplaceFunc: func(sub []string) string { return "[] " + sub[1] },
  },
  {
    Pattern:     regexp.MustCompile(`([aāiīuūṛṝḷḹeoēō])ḥ[^\w\n\r]*`),
    Explanation: "Word-final visarga is retained with marker [ḥV] to indicate echoing the vowel.",
    ReplaceFunc: func(sub []string) string { return sub[1] + "ḥ[" + sub[1] + "]" },
  },
}

type VisargaTransformer struct{}

func (t VisargaTransformer) Transform(input string) string {
	out, final := t.TransformWithExplanation(input)
  if (len(out) == 0) {
    return input
  } else {
		return final
	}
}

func (t VisargaTransformer) TransformWithExplanation(input string) ([]core.Transformation, string) {
  return core.ApplyRules(input, visargaRules)
}


// func applyRule(re *regexp.Regexp, text string, explanation string, replacement func([]string) string) (string, []sandhi.Transformation) {
//   transformations := []sandhi.Transformation{}
//
//   result := re.ReplaceAllStringFunc(text, func(match string) string {
//     sub := re.FindStringSubmatch(match)
//     if len(sub) < 2 {
//       return match
//     }
//     newStr := replacement(sub)
//     transformations = append(transformations, sandhi.Transformation{
//       Original:    match,
//       Transformed: newStr,
//       Explanation: explanation,
//     })
//     return newStr
//   })
//
//   return result, transformations
// }
// func (t VisargaTransformer) TransformWithExplanation(input string) ([]sandhi.Transformation, string) {
//   result := input
//   transformations := []sandhi.Transformation{}
//
//   for _, rule := range visargaRules {
//     result = rule.Pattern.ReplaceAllStringFunc(result, func(match string) string {
//       sub := rule.Pattern.FindStringSubmatch(match)
//       if len(sub) < 2 {
//         return match
//       }
//       newStr := rule.ReplaceFunc(sub, t.ChantingMode)
//       transformations = append(transformations, sandhi.Transformation{
//         Original:    match,
//         Transformed: newStr,
//         Explanation: rule.Explanation,
//       })
//       return newStr
//     })
//   }
//
//   return transformations, result
// }
//
// // re1 := regexp.MustCompile(`ḥ\s+(k|kh|p|ph|t|th)`)
// // var (
// //     visargaVowelRe         = regexp.MustCompile(`ḥ\s*([aāiīuūṛṝḷḹeoēō])`)
// //     visargaSoftConsonantRe = regexp.MustCompile(`ḥ\s*(k|kh|p|ph)`)
// //     visargaSibilantRe      = regexp.MustCompile(`ḥ\s*(ś|ṣ|s)`)
// //     visargaFinalRe         = regexp.MustCompile(`([aāiīuūṛṝḷḹeoēō])ḥ[^\w\n\r]*$`)
// // )
// //
// // func TransformVisarga(text string) string {
// //     text = visargaVowelRe.ReplaceAllString(text, "[o] $1")
// //     text = visargaSoftConsonantRe.ReplaceAllString(text, "[f] $1")
// //     text = visargaSibilantRe.ReplaceAllStringFunc(text, func(match string) string {
// //         sub := visargaSibilantRe.FindStringSubmatch(match)
// //         if (len(sub) < 2) {
// //             return match
// //         }
// //         sib := sub[1]
// //         return "[" + sib + sib + "]"
// //     })
// //     text = visargaFinalRe.ReplaceAllString(text, "$1[ḥ$1]")
// //     return text
// // }
// // func (t VisargaTransformer) TransformWithExplanation(input string) ([]sandhi.Transformation, string) {
// //     transformations := []sandhi.Transformation{}
// //     result := input
// //
// // 		// Rule 1: ḥ → f (upadhmānīya) before voiceless bilabial stops
// //     re1 := regexp.MustCompile(`ḥ\s+(p|ph)`)
// //     explanation := "‘ḥ’ becomes ‘f’ (upadhmānīya) before voiceless bilabial stops (p/ph)"
// //     result, tf := applyRule(re1, result, explanation, func(sub []string) string {
// //       return "[f] " + sub[1]
// //     })
// //     transformations = append(transformations, tf...)
// //
// // 		// Rule 2: ḥ → f (jihvāmūlīya) before voiceless velar stops
// //     re2 := regexp.MustCompile(`ḥ\s+(k|kh)`)
// //     explanation := "‘ḥ’ becomes ‘H’ (jihvāmūlīya) before voiceless velar stops (k/kh)"
// //     result, tf := applyRule(re2, result, explanation, func(sub []string) string {
// //       return "[H] " + sub[1]
// //     })
// //     transformations = append(transformations, tf...)
// //
// //     // Rule 3: ḥ → ō before vowels
// //     re3 := regexp.MustCompile(`ḥ\s+([aāiīuūṛṝḷḹeoēō])`)
// //     explanation := "‘ḥ’ becomes ‘ō ’ before a vowel."
// //     result, tf := applyRule(re3, result, explanation, func(sub []string) string {
// //       return "[ō] " + sub[1]
// //     })
// //     transformations = append(transformations, tf...)
// //
// //     // Rule 4: ḥ → removed before sibilants
// //     re4 := regexp.MustCompile(`ḥ\s+([śṣs])`)
// //     explanation := "‘ḥ’ is dropped before sibilants like ‘ś’, ‘ṣ’, or ‘s’."
// //     result, tf := applyRule(re4, result, explanation, func(sub []string) string {
// //       return "[] " + sub[1]
// //     })
// //     transformations = append(transformations, tf...)
// //
// // 		// Rule 5: Word-final visarga
// //     re5 := regexp.MustCompile(`([aāiīuūṛṝḷḹeoēō])ḥ[^\w\n\r]*$`)
// // 		explanation := "Word-final visarga is retained with contextual marker [ḥV] to indicate echoing the preceding vowel."
// //     result, tf := applyRule(re5, result, explanation, func(sub []string) string {
// //       return sub[1] + "[ḥ" + sub[1] + "]"
// //     })
// //     transformations = append(transformations, tf...)
// //
// //     return transformations, result
// // }
