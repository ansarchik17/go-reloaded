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

		// Markup processing: (cap), (low), (up)
		for i := 0; i < len(words); i++ {
			val := words[i]

			// (cap) 
			if len(val) >= 5 && strings.HasPrefix(val, "(cap") {
				if val == "(cap)" && i > 0 {
					Cap(words, 1, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasSuffix(val, ")") {
					k := TakeNumFromString(words[i])
					Cap(words, k, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasSuffix(words[i+1], ")") && val == "(cap," && i > 0 {
					k := TakeNumFromString(words[i+1])
					Cap(words, k, i-1)
					words = append(words[:i], words[i+2:]...)
					i -= 2
				} else if val == "(cap" && words[i+1] == ")" {
					Cap(words, 1, i-1)
					words = append(words[:i], words[i+2:]...)
					i -= 2
				}

				// (low)
			} else if len(val) >= 5 && strings.HasPrefix(val, "(low") {
				if val == "(low)" && i > 0 {
					Low(words, 1, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasSuffix(val, ")") {
					k := TakeNumFromString(words[i])
					Low(words, k, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasSuffix(words[i+1], ")") && val == "(low," && i > 0 {
					k := TakeNumFromString(words[i+1])
					Low(words, k, i-1)
					words = append(words[:i], words[i+2:]...)
					i -= 2
				}

				// (up)
			} else if len(val) >= 4 && strings.HasPrefix(val, "(up") {
				if val == "(up)" && i > 0 {
					Up(words, 1, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasSuffix(val, ")") {
					k := TakeNumFromString(words[i])
					Up(words, k, i-1)
					words = append(words[:i], words[i+1:]...)
					i--
				} else if strings.HasSuffix(words[i+1], ")") && val == "(up," && i > 0 {
					k := TakeNumFromString(words[i+1])
					Up(words, k, i-1)
					words = append(words[:i], words[i+2:]...)
					i -= 2
				}
			}
		}

		// Punctuation handling (, ; ! ? : . ...)
		for i := 0; i < len(words); i++ {
			val := words[i]

			if (val == "," || val == ";" || val == "!" || val == "?" || val == ":" || val == ".") && i > 0 {
				words[i-1] += val
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}

			if i > 0 && len(val) > 1 && strings.Contains(",;!?@.:", string(val[0])) {
				words[i-1] += string(val[0])
				words[i] = val[1:]
			}

			if len(val) >= 3 && val[:3] == "..." && i > 0 {
				words = apnd(words, "...", i)
				if len(val) == 3 {
					words = append(words[:i], words[i+1:]...)
					i--
				} else {
					words[i] = val[3:]
				}
			}

			if len(val) >= 2 && val[:2] == "!?" && i > 0 {
				words[i-1] += "!?"
				if len(val) == 2 {
					words = append(words[:i], words[i+1:]...)
					i--
				} else {
					words[i] = val[2:]
				}
			}
		}

		words = CleanStr(words)
		words = punc(words)
		words = fend(words)
		
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
				s[m] = strings.ToUpper(s[m][:1]) + s[m][1:]
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

func HexToDec(s string) string {
	s = strings.TrimSpace(s)
	temp, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s
	}
	res := strconv.Itoa(int(temp))
	return res
}

func BinToDec(s string) string {
	s = strings.TrimSpace(s)
	temp, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s
	}
	res := strconv.Itoa(int(temp))
	return res
}

func CleanStr(s []string) []string {
	ss := ""
	for i := 0; i < len(s); i++ {
		val := s[i]
		if val != "" {
			ss += val
			if i < len(s)-1 {
				ss += " "
			}
		}
	}
	return strings.Fields(ss)
}

func isWord(s string) bool {
	for _, val := range s {
		if strings.Contains(alp, strings.ToLower(string(val))) {
			return true
		}
	}
	return false
}

func notWord(s string) bool {
	for _, val := range s {
		if strings.Contains(alp, strings.ToLower(string(val))) {
			return false
		}
	}
	return true
}

func apnd(words []string, s string, n int) []string {
	for i := n - 1; i >= 0; i-- {
		if isWord(words[i]) {
			words[i] += s
			break
		}
	}
	return words
}

func punc(words []string) []string {
	for i := 0; i < len(words); i++ {
		val := words[i]
		if (val == "," || val == ";" || val == "!" || val == "?" || val == ":" || val == ".") && i > 0 {
			words[i-1] += val
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

func fend(words []string) []string {
	for i := 0; i < len(words); i++ {
		if notWord(words[i]) && i > 0 {
			words[i-1] += words[i]
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}
