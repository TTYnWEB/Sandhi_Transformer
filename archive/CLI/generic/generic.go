package generic

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Function to convert IAST uppercase characters to lowercase
func toLowerIAST(input string) string {
  // Iterate over the string and convert each character
  var result strings.Builder
  for _, r := range input {
  	// If the character is an uppercase IAST letter, convert it to lowercase
  	if unicode.IsUpper(r) {
  		result.WriteRune(unicode.ToLower(r))
  	} else {
  		result.WriteRune(r)
  	}
  }
  return result.String()
}

// Function to remove non-IAST characters from a string (only lowercase IAST)
func removeNonIASTChars(input string) string {
	// Regular expression pattern to match only lowercase IAST characters
	// Matches: lowercase letters, vowels with diacritics, anusvara (ṁ), visarga (ḥ), and other diacritic characters
	re := regexp.MustCompile(`[^a-zāīūōēṛṝḷḹṁḥṅñṭṇśṣ]`)

	// Replace all non-IAST characters with an empty string
	return re.ReplaceAllString(input, "")
}

func genericTransforms(s string) string {
	s = toLowerIAST(s)
	s = strings.ReplaceAll(s, "ṃ", "ṁ")
  s = strings.ReplaceAll(s, "o", "ō")
  s = strings.ReplaceAll(s, "e", "ē")
  s = strings.ReplaceAll(s, ":", "ḥ")
	s = removeNonIASTChars(s)
  return s
}
