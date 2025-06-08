package main

import (
  "fmt"
  "os"
  "sandhi_transformer/sandhi"
)

func usage() {
  fmt.Println("Usage: go run main.go <inputfile>")
  os.Exit(1)
}

func printErrNQuit(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	os.Exit(1)
}

func main() {
  if (len(os.Args) != 2) { usage() }

  inputPath := os.Args[1]
  
  data, err := os.ReadFile(inputPath)
  if (err != nil) { printErrNQuit("Failed to read input file", err) }
  
  transformed := sandhi.ApplyAll(string(data))
  
  outputPath := inputPath[:len(inputPath)-4] + "-transformed.txt"
  if err := os.WriteFile(outputPath, []byte(transformed), 0644); (err != nil) {
		printErrNQuit("Failed to write output file", err)
	}
  
  fmt.Printf("Transformed output written to: %s\n", outputPath)
}
