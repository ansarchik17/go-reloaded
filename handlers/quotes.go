package handlers

import "strings"

// Merge quotes into surrounding words safely
func MergeQuotes(words []string) []string {
	count := 0
	for _, val := range words {
		if val == "'" || strings.HasPrefix(val, "'") || strings.HasSuffix(val, "'") {
			count++
		}
	}
	count /= 2

	for i := 0; i < len(words); i++ {
		val := words[i]

		if strings.HasPrefix(val, "'") || strings.HasSuffix(val, "'") {
			if count > 0 {
				if val == "'" {
					if i < len(words)-1 {
						words[i+1] = val + words[i+1]
						words = append(words[:i], words[i+1:]...)
						i--
						count--
					}
				} else if strings.HasSuffix(val, "'") {
					words[i] = words[i][:len(words[i])-1]
					words[i+1] = val[len(val)-1:] + words[i+1]

					count--
				} else if strings.HasPrefix(val, "'") {
					count--
				}
			}
		}
	}
	return words
}

// Merge quotes into surrounding words safely
func MergeDQuotes(words []string) []string {
	count := 0
	for _, val := range words {
		if val == "\"" || strings.HasPrefix(val, "\"") || strings.HasSuffix(val, "\"") {
			count++
		}
	}
	count /= 2

	for i := 0; i < len(words); i++ {
		val := words[i]

		if strings.HasPrefix(val, "\"") || strings.HasSuffix(val, "\"") {
			if count > 0 {
				if val == "\"" {
					if i < len(words)-1 {
						words[i+1] = val + words[i+1]
						words = append(words[:i], words[i+1:]...)
						i--
						count--
					}
				} else if strings.HasSuffix(val, "\"") {
					words[i] = words[i][:len(words[i])-1]
					words[i+1] = val[len(val)-1:] + words[i+1]

					count--
				} else if strings.HasPrefix(val, "\"") {
					count--
				}
			}
		}
	}
	return words
}
