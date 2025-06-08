package core

func JoinSandhi(left, right, insert string, dropLeft, dropRight int) string {
	leftRunes := []rune(left)
	rightRunes := []rune(right)

	leftHead := string(leftRunes[:len(leftRunes)-dropLeft])
	rightTail := string(rightRunes[dropRight:])

	return leftHead + insert + rightTail
}

func ApplySandhiRules(input string, rules []SandhiRule) []Chunk {
	for _, rule := range rules {
		loc := rule.Pattern.FindStringSubmatchIndex(input)
		if loc == nil {
			continue
		}

		leftContext := input[:loc[0]]
		match := input[loc[0]:loc[1]]
		rightContext := input[loc[1]:]

		groups := rule.Pattern.FindStringSubmatch(match)
		if len(groups) < 3 {
			return []Chunk{{Text: input}} // Pattern mismatch
		}

		// Apply transformation on the matched portion
		transformed := rule.ReplaceFunc(match, groups[1], groups[2])

		return []Chunk{
			{Text: leftContext},
			{Text: transformed, Tooltip: rule.Explanation},
			{Text: rightContext},
		}
	}

	return []Chunk{{Text: input}} // No match
}
