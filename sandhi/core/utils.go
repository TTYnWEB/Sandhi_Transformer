package core

func ApplyRules(input string, rules []SandhiRule) ([]Transformation, string) {
  result := input
  transformations := []Transformation{}

  for _, rule := range rules {
    result = rule.Pattern.ReplaceAllStringFunc(result, func(match string) string {
      sub := rule.Pattern.FindStringSubmatch(match)
      if len(sub) < 1 {
        return match
      }
      newStr := rule.ReplaceFunc(sub)
      transformations = append(transformations, Transformation{
        Original:    match,
        Transformed: newStr,
        Explanation: rule.Explanation,
      })
      return newStr
    })
  }

  return transformations, result
}
