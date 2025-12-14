package handlers

import (
	"strconv"
	"strings"
)

// Ucl processes (cap), (low), (up) commands
func Ucl(words []string) []string {
	i := 0
	for i < len(words) {
		val := words[i]

		// (cap)
		if strings.HasPrefix(val, "(cap") {
			processed := processCommand(words, i, "cap", Cap)
			if processed > 0 {
				words = append(words[:i], words[i+processed:]...)
				if i > 0 {
					i--
				}
			} else {
				i++
			}
		} else if strings.HasPrefix(val, "(low") {
			// (low)
			processed := processCommand(words, i, "low", Low)
			if processed > 0 {
				words = append(words[:i], words[i+processed:]...)
				if i > 0 {
					i--
				}
			} else {
				i++
			}
		} else if strings.HasPrefix(val, "(up") {
			// (up)
			processed := processCommand(words, i, "up", Up)
			if processed > 0 {
				words = append(words[:i], words[i+processed:]...)
				if i > 0 {
					i--
				}
			} else {
				i++
			}
		} else {
			i++
		}
	}

	return words
}

// processCommand handles a command like (cap), (cap, 6), etc.
// Returns the number of words consumed by the command
func processCommand(words []string, idx int, cmdType string, fn func([]string, int, int)) int {
	val := words[idx]

	// Handle case where command is at start (no previous word)
	// Still need to consume the command and any following parts
	if idx == 0 {
		// Check if it's a simple command like (cap)
		if val == "("+cmdType+")" {
			return 1
		}
		// Check if it's split like (cap, )
		if val == "("+cmdType+"," || strings.HasPrefix(val, "("+cmdType+",") {
			if idx+1 < len(words) {
				nextWord := words[idx+1]
				// If next word is ")", consume both
				if nextWord == ")" || (strings.HasPrefix(nextWord, ")") && len(strings.TrimPrefix(nextWord, ")")) == 0) {
					return 2
				}
				// Might have a number, check
				if idx+2 < len(words) && words[idx+2] == ")" {
					return 3
				}
				if strings.HasSuffix(nextWord, ")") {
					return 2
				}
			}
			return 1
		}
		// Complete command with parameter
		if strings.HasSuffix(val, ")") && strings.Contains(val, ",") {
			return 1
		}
		return 1
	}

	// Simple case: (cap), (low), (up)
	if val == "("+cmdType+")" {
		fn(words, 1, idx-1)
		return 1
	}

	// Check if it's a complete command like (cap,6) or (cap, 6)
	if strings.HasSuffix(val, ")") && strings.Contains(val, ",") {
		// Check for negative numbers
		if strings.Contains(val, "-") {
			// Invalid (negative), just remove it
			return 1
		}
		k := TakeNumFromString(val)
		if k > 0 {
			// Cap at the number of available words
			if k > idx {
				k = idx
			}
			if k > 0 {
				fn(words, k, idx-1)
				return 1
			}
		}
		// Invalid parameter, default to 1 word
		if idx > 0 {
			fn(words, 1, idx-1)
		}
		return 1
	}

	// Check if command is split: (cap, or (cap,
	if val == "("+cmdType+"," || strings.HasPrefix(val, "("+cmdType+",") {
		// Look for the number and closing paren in following words
		// Check next word(s) for number and closing paren
		if idx+1 < len(words) {
			nextWord := words[idx+1]

			// Case 1: next word is just the number, and word after is ")"
			if idx+2 < len(words) && words[idx+2] == ")" {
				// Check for negative
				if strings.Contains(nextWord, "-") {
					// Invalid (negative), remove command
					return 3
				}
				k := TakeNumFromString(nextWord)
				if k > 0 {
					// Cap at available words
					if k > idx {
						k = idx
					}
					if k > 0 {
						fn(words, k, idx-1)
						return 3 // consumed: (cap,, number, )
					}
				}
				// Invalid parameter, default to 1 word
				if idx > 0 {
					fn(words, 1, idx-1)
				}
				return 3
			}

			// Case 2: next word contains number and closing paren like "6)" or "a)"
			if strings.HasSuffix(nextWord, ")") {
				// Check for negative
				if strings.Contains(nextWord, "-") {
					// Invalid (negative), remove command
					return 2
				}
				k := TakeNumFromString(nextWord)
				if k > 0 {
					// Cap at available words
					if k > idx {
						k = idx
					}
					if k > 0 {
						fn(words, k, idx-1)
						return 2 // consumed: (cap,, number)
					}
				}
				// Invalid parameter (like "a)"), default to 1 word
				if idx > 0 {
					fn(words, 1, idx-1)
				}
				return 2
			}

			// Case 3: next word is just a number (might have trailing punctuation)
			// Check for negative
			if !strings.Contains(nextWord, "-") {
				k := TakeNumFromString(nextWord)
				if k > 0 {
					// Look for closing paren in next words
					if idx+2 < len(words) {
						if words[idx+2] == ")" || strings.HasPrefix(words[idx+2], ")") {
							// Cap at available words
							if k > idx {
								k = idx
							}
							if k > 0 {
								fn(words, k, idx-1)
								if strings.HasPrefix(words[idx+2], ")") && len(words[idx+2]) > 1 {
									return 2
								}
								return 3
							}
						}
					}
				}
			}

			// Case 4: next word is ")" or starts with ")" - handle (up, )
			if nextWord == ")" || (strings.HasPrefix(nextWord, ")") && len(strings.TrimPrefix(nextWord, ")")) == 0) {
				// Invalid format (no number), default to 1 word if we have a previous word
				if idx > 0 {
					fn(words, 1, idx-1)
				}
				// Always consume both (up, and )
				return 2
			}

			// Case 5: next word is not a number and not ")"
			// Might be like (up, a) - default to 1 word
			if idx > 0 {
				fn(words, 1, idx-1)
			}
			// Try to find and consume closing paren
			if idx+2 < len(words) && (words[idx+2] == ")" || strings.HasPrefix(words[idx+2], ")")) {
				return 3
			}
			return 2
		}

		// Invalid format, but if we can, default to 1 word
		if idx > 0 {
			fn(words, 1, idx-1)
		}
		return 1
	}

	// Not a recognized command format, don't consume
	return 0
}

func Cap(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if IsWord(s[m]) {
			tp := FindFl(s[m])
			if len(s[m]) == 1 {
				s[m] = strings.ToUpper(s[m])
			} else {
				s[m] = strings.ToUpper(s[m][:tp+1]) + strings.ToLower(s[m][tp+1:])
			}
			n--
		}
		m--
	}
}

func Low(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if IsWord(s[m]) {
			s[m] = strings.ToLower(s[m])
			n--
		}
		m--
	}
}

func Up(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if IsWord(s[m]) {
			s[m] = strings.ToUpper(s[m])
			n--
		}
		m--
	}
}

func TakeNumFromString(s string) int {
	res := ""
	for _, val := range s {
		if val >= '0' && val <= '9' {
			res += string(val)
		}
	}
	if res == "" {
		return 0
	}
	ans, _ := strconv.Atoi(res)
	return ans
}

func FindFl(s string) int {
	for i, val := range s {
		if val >= 'a' && val <= 'z' || val >= 'A' && val <= 'Z' {
			return i
		}
	}
	return -1
}

//a