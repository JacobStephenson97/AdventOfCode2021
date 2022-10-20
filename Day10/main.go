package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		var buff []rune
		for j, char := range line {
			buff = append(buff, char)
			if ']' == char {
				if checkBuff(buff, char) {
					println(i, j)
				}
			} else if ')' == char {
			} else if '}' == char {
			} else if '>' == char {
			}
		}
	}
}

func checkBuff(buff []rune, char rune) bool {
	dict := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
		'>': '<',
	}
	for i := len(buff) - 1; i >= 0; i-- {
    if buff[i-1] == ''
		else if buff[i-1] == dict[char] {
			return true
		} else {
			return false
		}
	}
}
