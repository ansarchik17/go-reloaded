package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const alp = "abcdefghijklmnopqrstuvwxyz"

func main() {

	fmt.Println()
	if len(os.Args) != 3{
		fmt.Println("Error ! ")
		return 
	}
	readF := os.Args[1]
	writeF := os.Args[2]
	content, _ := os.ReadFile(readF)

	words := strings.Split(string(content), " ")

	// ss := words[1]
	// fmt.Println(ss[:5])


	// to CAP and LOW and UP :

	for i := 0; i < len(words); i++ {
		val := words[i]

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
			}else if strings.HasSuffix(words[i+1], ")") && val == "(cap," && i > 0 {
				k := TakeNumFromString(words[i+1])
				Cap(words, k, i-1)
				words = append(words[:i], words[i+2:]...)
				i -= 2
			} else if val == "(cap" && words[i+1] == ")" {
				Cap(words, 1, i-1)
				words = append(words[:i], words[i+2:]...)
				i -= 2
			}
		} else if len(val) >= 5 && strings.HasPrefix(val, "(low") {
			if val == "(low)" && i > 0 {
				Low(words, 1, i-1)
				words = append(words[:i], words[i+1:]...)
				i--
			} else if strings.HasSuffix(words[i+1], ")") && val == "(low," && i > 0 {
				k := TakeNumFromString(words[i+1])
				Low(words, k, i-1)
				words = append(words[:i], words[i+2:]...)
				i -= 2
			}
		} else if len(val) >= 4 && strings.HasPrefix(val, "(up") {
			if val == "(up)" && i > 0 {
				Up(words, 1, i-1)
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
	
	// COMMA - TOMMA

	

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
	
	

	words = CleanStr(words)
	
	for i := 0; i < len(words); i++ {
		val := words[i]
		
		if val == "," || val == ";" || val == "!" || val == "?" || val == ":" || val == "." {

			words[i-1] += val
			if i == len(words)-1 && i > 0 {
				words = words[:i]
				break
			} else {
				if i > 0 {
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}

		}

		if i > 0 && len(val) > 1 && (val[0] == ',' || val[0] == ';' || val[0] == '!' || val[0] == '?' || val[0] == ':' || val == ".") {
			
			if val[0] != '!' && val[1] != '?'{

			words[i-1] += string(val[0])
			words[i] = val[1:]
			}


		}

		if len(val) >= 3 && val[:3] == "..." {
			if len(val) == 2 {
				words = apnd(words, "...", i)
				words = append(words[:i], words[i+1:]...)
				i--
			} else {
				words = apnd(words, "...", i)
				words[i] = val[3:]
			}
		}

		if len(val) >= 2 && val[:2] == "!?" {
			if len(val) == 2 {
				words[i-1] += "!?"
			
				words = append(words[:i], words[i+1:]...)
				
				i--
			} else {
				words[i-1] += val[:2]
				words[i] = val[2:]
			}
		}

	}

	

	vowels := "aeiou"
	// HEX and BIN and An and an
	for i := 0; i < len(words); i++ {
		val := words[i]

		if val == "(hex)" && i != 0 {
			words[i-1] = HexToDec(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}

		if val == "(bin)" && i != 0 {
			words[i-1] = BinToDec(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	

	for i := 0; i < len(words); i++ {
		val := words[i]

		if val == "a" || val == "A" {
			if i != len(words)-1 && words[i+1] != "" {
				if words[i+1] == "and" {
					continue
				} else if strings.Contains(vowels, strings.ToLower(string(words[i+1][0]))) && len(words[i+1]) > 1 {
					if val == "a" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
				}
			}
		}

		if val == "an" || val == "An" {
			if i != len(words)-1 && words[i+1] != "" {
				if words[i+1] == "and" {
					continue
				} else if !strings.Contains(vowels, strings.ToLower(string(words[i+1][0]))) && len(words[i+1]) > 1 {
					if val == "an" {
						words[i] = "a"
					} else {
						words[i] = "A"
					}
				}
			}
		}

	}
	

	fmt.Printf("Initially: %v\n", string(content))

	fmt.Println()

	contPaste := strings.Join(words, " ")
	// contPaste = strings.TrimRight(contPaste, "!")
	os.WriteFile(writeF, []byte(contPaste), 0o644)

	contR, _ := os.ReadFile(writeF)
	fmt.Printf("Result: %v\n", string(contR))

	fmt.Println()
}

func Cap(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if s[m] != "" {
			if len(s[m]) == 1 {
				s[m] = strings.ToUpper(s[m])
				n--
			} else {
				s[m] = strings.ToUpper(s[m][:1]) + s[m][1:]
				n--
			}
		}

		m--

	}
}

func Low(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if s[m] != "" {
			s[m] = strings.ToLower(s[m])
			n--
		}

		m--

	}
}

func Up(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if s[m] != "" {
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
		} else {
			break
		}
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
			if val == ")" && i != 0 {
				ss = ss[:len(ss)-1]
				ss += val
			} else {
				ss += val
			}

			if i < len(s)-1 {
				ss += " "
			}
		}

	}
	fmt.Println(ss)
	return strings.Split(ss, " ")
}

func isWord(s string) bool {
	for _, val := range s {
		if strings.Contains(alp, string(val)) {
			return true
		}
	}
	return false
}

func apnd(words []string, s string, n int) []string{

	for i :=n-1; i >=0; i--{
		if isWord(words[i]){
			words[i] += s
			break
		}
	}
	return words
}