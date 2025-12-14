package main

import (
	"fmt"
	"go-reloaded/funcs"
	"os"
	"strings"
)

const alp = "abcdefghijklmnopqrstuvwxyz"

func main() {

	fmt.Println()
	if len(os.Args) != 3 {
		fmt.Println("Error ! ")
		return
	}
	readF := os.Args[1]
	writeF := os.Args[2]

	content, _ := os.ReadFile(readF)

	// Handle multiple lines correctly
	lines := strings.Split(string(content), "\n")
	for li, line := range lines {
		if line == "" {
			continue
		}
		words := strings.Fields(line)

		// Preprocess punctuation: separate leading/trailing punctuation
		words = funcs.SeparatePunc(words)

		// Process hex/bin conversions first (they work on the word before)
		words = funcs.ProcessHexBin(words)

		// Process case commands: (cap), (low), (up)
		words = funcs.Ucl(words)

		// Reattach punctuation after case changes
		words = funcs.ReattachPunc(words)

		// Merge quotes
		words = funcs.MergeQuotes(words)
		words = funcs.MergeDQuotes(words)

		// Fix articles (a/an)
		words = funcs.FixArticles(words)

		lines[li] = strings.Join(words, " ")
	}

	contPaste := strings.Join(lines, "\n")

	os.WriteFile(writeF, []byte(contPaste), 0o644)

	contR, _ := os.ReadFile(writeF)
	fmt.Printf("Initially: %v\n", string(content))
	fmt.Println()
	fmt.Printf("Result: %v\n", string(contR))
	fmt.Println()
}
