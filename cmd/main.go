package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"path/filepath"
	"sandhi_transformer/sandhi"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  go run main.go <inputfile>")
	fmt.Println("  OR")
	fmt.Println("  go run main.go -i   (interactive line input)")
	os.Exit(1)
}

func printErrNQuit(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	os.Exit(1)
}

func runInteractive() {
	fmt.Print("Enter text to transform: ")
	reader := bufio.NewReader(os.Stdin)
	inputText, err := reader.ReadString('\n')
	if (err != nil) { printErrNQuit("Failed to read input", err) }

	transformations, result := sandhi.ApplyAllSandhiWithExplanation(inputText)
	fmt.Println("Transformed:", result)
  for _, t := range transformations {
    fmt.Printf("- '%s' → '%s': %s\n", t.Original, t.Transformed, t.Explanation)
  }
}

func runFileMode(filePath string) {
	data, err := os.ReadFile(filePath)
	if (err != nil) { printErrNQuit("Failed to read input file", err) }
	transformed := sandhi.ApplyAllSandhi(string(data))

	ext := filepath.Ext(filePath)
	base := strings.TrimSuffix(filePath, ext)
	outputPath := base + "-transformed.txt"

	err = os.WriteFile(outputPath, []byte(transformed), 0644)
	if (err != nil) { printErrNQuit("Failed to write output file", err) }
	fmt.Printf("Transformed output written to: %s\n", outputPath)
}

func main() {
	if ((len(os.Args)) != 2) { usage() }

	switch arg := os.Args[1]; arg {
	  case "-i":
	  	runInteractive()
	  default:
	  	runFileMode(arg)
	}
}

// transformed := sandhi.ApplyAll(strings.TrimSpace(inputText))
// fmt.Println("Transformed text:", transformed)
