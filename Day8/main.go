package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	parsedInput := parseInput(input)

	count := 0
	for _, line := range parsedInput {
		outputs := strings.Split(line[1], " ")
		for _, output := range outputs {
			if len(output) == 2 || len(output) == 4 || len(output) == 7 || len(output) == 3 {
				count++
			}
		}
	}
	println(count)
	partTwo(parsedInput)
}

func partTwo(parsedInput [][]string) {
	// 0000
	//1    2
	//1    2
	// 3333
	//4    5
	//4    5
	// 6666

	// identify 2 and 5

	var finalNumbers []string
	for _, line := range parsedInput {
		layout := make([]string, 7)
		inputs := strings.Split(line[0], " ")
		for _, input := range inputs {
			if len(input) == 2 {
				layout[2] = string(input[0])
				layout[5] = string(input[1])
			}
		}
		// identify 0
		for _, input := range inputs {
			if len(input) == 3 {
				for _, char := range input {
					if string(char) != layout[2] && string(char) != layout[5] {
						layout[0] = string(char)
					}
				}
			}
		}
		//indentify 1 and 3
		for _, input := range inputs {
			if len(input) == 4 {
				for _, char := range input {
					if string(char) != layout[0] && string(char) != layout[2] && string(char) != layout[5] {
						for _, input2 := range inputs {
							if len(input2) == 6 && (!strings.Contains(input2, string(char)) && layout[2] != string(char)) {
								layout[3] = string(char)
							}
						}
					}
				}
			}
		}
		for _, input := range inputs {
			if len(input) == 4 {
				for _, char := range input {
					if string(char) != layout[1] && string(char) != layout[2] && string(char) != layout[3] && string(char) != layout[5] {
						layout[1] = string(char)
					}
				}
			}
		}
		//identify 4 and 6
		for _, input := range inputs {
			if len(input) == 6 {
				if strings.Contains(input, layout[2]) && strings.Contains(input, layout[5]) && strings.Contains(input, layout[3]) {
					for _, char := range input {
						if !contains(layout, string(char)) {
							layout[6] = string(char)
						}
					}
				}
			}
		}

		for _, input := range inputs {
			if len(input) == 7 {
				for _, char := range input {
					if !contains(layout, string(char)) {
						layout[4] = string(char)
					}
				}
			}
		}
		outputs := strings.Split(line[1], " ")
		numberstring := ""
		for _, output := range outputs {
			length := len(output)
			switch {
			case length == 2:
				numberstring += "1"
			case length == 3:
				numberstring += "7"
			case length == 4:
				numberstring += "4"
			case length == 7:
				numberstring += "8"
			case length == 5:
				if strings.Contains(output, layout[2]) && strings.Contains(output, layout[5]) {
					numberstring += "3"
				} else if strings.Contains(output, layout[1]) {
					numberstring += "5"
				} else {
					numberstring += "2"
				}
			case length == 6:
				if strings.Contains(output, layout[2]) && strings.Contains(output, layout[5]) && strings.Contains(output, layout[3]) {
					numberstring += "9"
				} else if strings.Contains(output, layout[2]) && strings.Contains(output, layout[5]) {
					numberstring += "0"
				} else {
					numberstring += "6"
				}
			}
		}
		finalNumbers = append(finalNumbers, numberstring)
	}
	println(calcTotal(finalNumbers))
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\r")
	parsedInput := make([][]string, len(lines))
	for i, line := range lines {
		parsedInput[i] = []string{strings.TrimSpace(strings.Split(line, "|")[0]), strings.TrimSpace(strings.Split(line, "|")[1])}
	}
	return parsedInput
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func check(number string, output string) bool {
	for _, char := range output {
		if !strings.Contains(number, string(char)) {
			return false
		}
	}
	return true
}

func calcTotal(numbers []string) int {
	total := 0
	for _, number := range numbers {
		nummy, _ := strconv.Atoi(number)
		total += nummy
	}
	return total
}
