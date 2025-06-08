package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

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

// const vowels = `[aāiīuūṛṝḷḹeo]`
// var consonantSandhiRules = []struct {
// 	Pattern *regexp.Regexp
// 	Replace string
// }{
// 	{regexp.MustCompile(`k\s+k`), "kk"},
// 	{regexp.MustCompile(`t\s+t`), "tt"},
// 	{regexp.MustCompile(`d\s+d`), "dd"},
// 	{regexp.MustCompile(`p\s+p`), "pp"},
// 	{regexp.MustCompile(`t\s+c`), "cc"},
// 	{regexp.MustCompile(`t\s+ch`), "cch"},
// 	{regexp.MustCompile(`d\s+dh`), "ddh"},
// 	{regexp.MustCompile(`n\s+c`), "ñc"},
// 	{regexp.MustCompile(`n\s+ch`), "ñch"},
// 	{regexp.MustCompile(`n\s+j`), "ñj"},
// 	{regexp.MustCompile(`n\s+jh`), "ñjh"},
// 	{regexp.MustCompile(`s\s+k`), "sk"},
// 	{regexp.MustCompile(`s\s+p`), "sp"},
// 	{regexp.MustCompile(`s\s+t`), "st"},
// 	{regexp.MustCompile(`s\s+c`), "sc"},
// 	{regexp.MustCompile(`s\s+([aāiīuūṛeo])`), "r $1"}, // optional
// }

var (
	// Anusvāra
	anusvaraRe              = regexp.MustCompile(`ṁ(\s*)([ckgjṭḍtdpbṅñṇnm])(h?)`)
	// Visarga
	visargaVowelRe          = regexp.MustCompile(`ḥ\s*([aāiīuūṛeo])`)
	visargaSoftConsonantRe  = regexp.MustCompile(`ḥ\s*(k|kh|p|ph)`)
	visargaSibilantRe       = regexp.MustCompile(`ḥ\s*(ś|ṣ|s)`)
	visargaFinalRe          = regexp.MustCompile(`([aāiīuūṛeēo])ḥ(\s|$|[.,;?!])`)
	// Consonant Sandhi: (basic patterns)
	consonantRules = []struct {
		pattern *regexp.Regexp
		replace string
	}{
		{regexp.MustCompile(`t\s+(ś)`), "[cch]"},
		{regexp.MustCompile(`d\s+(dh)`), "[ddh]"},
		{regexp.MustCompile(`n\s+(ch)`), "[ñch]"},
		{regexp.MustCompile(`s\s+(k)`), "[sk]"},
	}
)

func transformAnusvara(text string) string {
	return anusvaraRe.ReplaceAllStringFunc(text, func(match string) string {
		sub := anusvaraRe.FindStringSubmatch(match)
		if len(sub) < 4 {
			return match
		}
		space := sub[1]
		cons := sub[2]
		asp := sub[3]
		key := cons + asp
		if nasal, ok := nasalMap[key]; ok {
			return "[" + nasal + "]" + space + key
		}
		return match
	})
}

func transformVisarga(text string) string {
	text = visargaVowelRe.ReplaceAllString(text, "[o] $1")
	text = visargaSoftConsonantRe.ReplaceAllString(text, "[f] $1")
	text = visargaSibilantRe.ReplaceAllStringFunc(text, func(match string) string {
		sub := visargaSibilantRe.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		sib := sub[1]
		return "[" + sib + sib + "]"
	})
	text = visargaFinalRe.ReplaceAllString(text, "$1[ḥ$1]$2")
	return text
}

func transformConsonant(text string) string {
	for _, rule := range consonantRules {
		text = rule.pattern.ReplaceAllStringFunc(text, func(match string) string {
			return rule.replace
		})
	}
	return text
}

func transformSandhi(text string) string {
	text = transformAnusvara(text)
	text = transformVisarga(text)
	text = transformConsonant(text)
	return text
}

func writeOutputFile(inputPath, content string) (string, error) {
	ext := filepath.Ext(inputPath)
	base := strings.TrimSuffix(inputPath, ext)
	outPath := base + "-transformed.txt"
	err := os.WriteFile(outPath, []byte(content), 0644)
	return outPath, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run sandhi.go <inputfile>")
		return
	}
	inputPath := os.Args[1]
	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
		os.Exit(1)
	}

	transformed := transformSandhi(string(inputData))

	outputPath, err := writeOutputFile(inputPath, transformed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Transformed output written to: %s\n", outputPath)
}
