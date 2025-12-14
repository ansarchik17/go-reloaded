package funcs

import (
	"strings"
)

// FixArticles adjusts "a"/"an" based on the following word
func FixArticles(words []string) []string {
	vowels := "aeiouAEIOU"
	// list of special words where 'h' is silent
	silentH := map[string]bool{
		"hour":   true,
		"honest": true,
		"honor":  true,
		"heir":   true,
	}

	for i := 0; i < len(words)-1; i++ {
		val := words[i]
		next := words[i+1]

		if !IsWord(next) {
			continue
		}

		if next == "and" || next == "or" {
			continue
		}

		lowerNext := strings.ToLower(next)
		if strings.EqualFold(val, "a") || strings.EqualFold(val, "an") {
			needsAn := false
			firstChar := next[0]

			if strings.Contains(vowels, string(firstChar)) {
				needsAn = true
			} else if silentH[lowerNext] {
				needsAn = true
			} else {
				needsAn = false
			}

			// apply the change preserving capitalization
			if needsAn {
				if strings.EqualFold(val, "a") {
					if val == "a" {
						words[i] = "an"
					} else if val == "A" {
						words[i] = "An"
					} else if val == "AN" {
						words[i] = "AN"
					}
				}
			} else {
				if strings.EqualFold(val, "an") {
					if val == "an" {
						words[i] = "a"
					} else if val == "An" {
						words[i] = "A"
					} else if val == "AN" {
						words[i] = "A"
					}
				}
			}
		}
	}
	return words
}
