package handlers

import (
	"strconv"
)

// ProcessHexBin processes (hex) and (bin) commands
func ProcessHexBin(words []string) []string {
	for i := 0; i < len(words); i++ {
		val := words[i]

		// Handle (hex)
		if val == "(hex)" && i > 0 {
			prevWord := words[i-1]
			// Check if previous word is a valid hex number
			if IsHexNumber(prevWord) {
				decimal, err := strconv.ParseInt(prevWord, 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decimal, 10)
					words = append(words[:i], words[i+1:]...)
					i--
				} else {
					// Invalid hex, remove the (hex) marker
					words = append(words[:i], words[i+1:]...)
					i--
				}
			} else {
				// Previous word is not a valid hex number, remove the (hex) marker
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		// Handle (bin)
		if val == "(bin)" && i > 0 {
			prevWord := words[i-1]
			// Check if previous word is a valid binary number
			if IsBinNumber(prevWord) {
				decimal, err := strconv.ParseInt(prevWord, 2, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decimal, 10)
					words = append(words[:i], words[i+1:]...) // i+1
					i--
				} else {
					// Invalid bin, remove the (bin) marker
					words = append(words[:i], words[i+1:]...)
					i--
				}
			} else {
				// Previous word is not a valid binary number, remove the (bin) marker
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}
	return words
}

// IsHexNumber checks if a string is a valid hexadecimal number
func IsHexNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if !((char >= '0' && char <= '9') || (char >= 'A' && char <= 'F') || (char >= 'a' && char <= 'f')) {
			return false
		}
	}
	return true
}

// IsBinNumber checks if a string is a valid binary number
func IsBinNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}
