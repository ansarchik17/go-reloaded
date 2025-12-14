package funcs

import "strings"

const alp = "abcdefghijklmnopqrstuvwxyz"

// Separate leading/trailing punctuation from words, including quotes
// But preserve command markers like (cap), (hex), etc.
func SeparatePunc(words []string) []string {
	var res []string
	for _, w := range words {
		// Check if this is a command marker - don't separate it
		if strings.HasPrefix(w, "(") && (strings.Contains(w, "cap") || strings.Contains(w, "low") ||
			strings.Contains(w, "up") || strings.Contains(w, "hex") || strings.Contains(w, "bin")) {
			res = append(res, w)
			continue
		}

		prefix := ""
		suffix := ""
		for len(w) > 0 && strings.Contains(".,;:!?'\"", string(w[0])) { // added quotes
			prefix += string(w[0])
			w = w[1:]
		}
		for len(w) > 0 && strings.Contains(".,;:!?'\"", string(w[len(w)-1])) { // added quotes
			suffix = string(w[len(w)-1]) + suffix
			w = w[:len(w)-1]
		}
		if prefix != "" {
			res = append(res, prefix)
		}
		if w != "" {
			res = append(res, w)
		}
		if suffix != "" {
			res = append(res, suffix)
		}
	}
	return res
}

// Reattach punctuation after capitalization, including quotes
// Also merges groups of the same punctuation (e.g., ..., !!, ???)
func ReattachPunc(words []string) []string {
	if len(words) == 0 {
		return words
	}

	var res []string
	i := 0

	for i < len(words) {
		w := words[i]

		// If it's a word, add it
		if IsWord(w) {
			res = append(res, w)
			i++
			continue
		}

		// It's punctuation - collect consecutive punctuation of the same type
		punctuation := w
		i++

		// Merge consecutive punctuation of the same type (e.g., . . . -> ...)
		for i < len(words) && !IsWord(words[i]) {
			nextPunc := words[i]
			// Check if we should merge (same punctuation type)
			if len(punctuation) > 0 && len(nextPunc) > 0 {
				lastChar := punctuation[len(punctuation)-1]
				firstChar := nextPunc[0]
				if lastChar == firstChar {
					// Same type, merge them
					punctuation += nextPunc
					i++
				} else {
					// Different type, stop merging
					break
				}
			} else {
				break
			}
		}

		// Attach punctuation to the previous word if available
		if len(res) > 0 {
			res[len(res)-1] += punctuation
		} else {
			// No previous word, just add punctuation (shouldn't happen normally)
			res = append(res, punctuation)
		}
	}

	return res
}

func IsWord(s string) bool {
	for _, val := range s {
		if strings.Contains(alp, strings.ToLower(string(val))) {
			return true
		}
	}
	return false
}
