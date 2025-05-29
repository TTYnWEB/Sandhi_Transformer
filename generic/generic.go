package generic

import "strings"

func genericTransforms(s string) string {
  s = strings.ReplaceAll(s, "o", "ō")
  s = strings.ReplaceAll(s, "e", "ē")
  s = strings.ReplaceAll(s, ":’", "ḥ’")
  s = strings.ReplaceAll(s, ":", "ḥ")
  return s
}
