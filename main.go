package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
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

	// for hours it is like protection
	content = bytes.ReplaceAll(content, []byte(":"), []byte(":@"))

	// Handle multiple lines correctly
	lines := strings.Split(string(content), "\n")
	for li, line := range lines {
		words := strings.Fields(line)

		// Preprocess punctuation: separate leading/trailing punctuation
		words = SeparatePunc(words)
		
		// Markup processing: (cap), (low), (up)
		for i := 0; i < len(words); i++ {
			val := words[i]

			// (cap) 
			if strings.HasPrefix(val, "(cap") {
				if val == "(cap)" && i > 0 {
					Cap(words, 1, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasSuffix(val, ")") {
					k := TakeNumFromString(val)
					Cap(words, k, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasPrefix(val, "(cap,") && i > 0 {
					k := TakeNumFromString(val)
					Cap(words, k, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}

			// (low)
			if strings.HasPrefix(val, "(low") {
				if val == "(low)" && i > 0 {
					Low(words, 1, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasSuffix(val, ")") {
					k := TakeNumFromString(val)
					Low(words, k, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}

			// (up)
			if strings.HasPrefix(val, "(up") {
				if val == "(up)" && i > 0 {
					Up(words, 1, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasSuffix(val, ")") {
					k := TakeNumFromString(val)
					Up(words, k, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}
		}

		words = ReattachPunc(words)
		// Quote handling: merge `'` in correct positions
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
 						count--
					}
				} else if strings.HasSuffix(val, "'") {
					words[i] = words[i][:len(words[i])-1]
				
					words[i+1] = val[len(val)-1:] + words[i+1]
				
					count--
					i++
				} else if strings.HasPrefix(val, "'") {
					count--
				}
			} else {
				if val == "'" {
					words[i-1] = words[i-1] + val
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasPrefix(val, "'") {
					words[i-1] = words[i-1] + string(val[0])
					words[i] = words[i][1:]
 				}
			}
		}
	}
		lines[li] = strings.Join(words, " ")
	}

	contPaste := strings.Join(lines, "\n")

	// (restore hours) 
	contPaste = strings.ReplaceAll(contPaste, ":@" , ":")

	os.WriteFile(writeF, []byte(contPaste), 0o644)

	contR, _ := os.ReadFile(writeF)
	fmt.Printf("Initially: %v\n", string(content))
	fmt.Println()
	fmt.Printf("Result: %v\n", string(contR))
	fmt.Println()
}

func Cap(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if isWord(s[m]) {
			if len(s[m]) == 1 {
				s[m] = strings.ToUpper(s[m])
			} else {
				s[m] = strings.ToUpper(s[m][:1]) + strings.ToLower(s[m][1:])
			}
			n--
		}
		m--
	}
}

func Low(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if isWord(s[m]) {
			s[m] = strings.ToLower(s[m])
			n--
		}
		m--
	}
}

func Up(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if isWord(s[m]) {
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

func isWord(s string) bool {
	for _, val := range s {
		if strings.Contains(alp, strings.ToLower(string(val))) {
			return true
		}
	}
	return false
}

// Separate leading/trailing punctuation from words, including quotes
func SeparatePunc(words []string) []string {
	var res []string
	for _, w := range words {
		prefix := ""
		suffix := ""
		for len(w) > 0 && strings.Contains(".,;:!?\"'", string(w[0])) { // added quotes
			prefix += string(w[0])
			w = w[1:]
		}
		for len(w) > 0 && strings.Contains(".,;:!?\"'", string(w[len(w)-1])) { // added quotes
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
func ReattachPunc(words []string) []string {
	var res []string
	for i := 0; i < len(words); i++ {
		w := words[i]
		// only attach punctuation that is not a word
		if !isWord(w) && len(res) > 0 {
			res[len(res)-1] += w
		} else {
			res = append(res, w)
		}
	}
	return res
}
