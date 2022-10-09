package main

import (
	_ "embed"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	parsedInput := strings.Split(strings.TrimSpace(input), "\n")

	// gamma := findGamma(parsedInput)
	// epislon := invertBinary((findGamma(parsedInput)))
	// println(fmt.Sprintf("Epsilon: %d", binaryToDecimal(epislon)))
	// println(fmt.Sprintf("Gamma: %d", binaryToDecimal(gamma)))

	// answer := binaryToDecimal(gamma) * binaryToDecimal(epislon)

	// println(answer)
	println(partTwo(parsedInput))
}

func findGamma(input []string) string {
	//find gamma

	final := ""

	for index := 0; index < len(input[0])-1; index++ {
		oneCount := 0
		zeroCount := 0
		for _, line := range input {
			if line[index] == '1' {
				oneCount++
			} else {
				zeroCount++
			}
		}
		if oneCount > zeroCount {
			final += "1"
		} else {
			final += "0"
		}
	}
	return final
}

func invertBinary(input string) string {
	final := ""
	for _, char := range input {
		if char == '1' {
			final += "0"
		} else {
			final += "1"
		}
	}
	return final
}

func binaryToDecimal(input string) int {
	decimal := 0
	println(len(input))
	for index, char := range input {
		if char == '1' {
			decimal += int(math.Pow(2, float64(len(input)-index-1)))
		}
	}
	return decimal
}

func partTwo(input []string) int {
	println("Most", recurseFind(input, 0, "most"))
	println("Least", recurseFind(input, 0, "least"))

	oxy := recurseFind(input, 0, "most")
	scrubber := recurseFind(input, 0, "least")

	return binaryToDecimal(strings.TrimSpace(oxy)) * binaryToDecimal(strings.TrimSpace(scrubber))
}

func recurseFind(input []string, index int, option string) string {
	if len(input) == 1 {
		return input[0]
	}
	var commonBit string
	if option == "most" {
		commonBit = mostCommonBit(input, index)
	} else {
		commonBit = leastCommonBit(input, index)
	}

	newInput := []string{}
	for _, line := range input {
		if string(line[index]) == commonBit {
			newInput = append(newInput, line)
		}
	}
	return recurseFind(newInput, index+1, option)
}

func mostCommonBit(input []string, index int) string {
	oneCount := 0
	zeroCount := 0
	for _, line := range input {
		if line[index] == '1' {
			oneCount++
		} else {
			zeroCount++
		}
	}
	if oneCount == zeroCount {
		return "1"
	}
	if oneCount > zeroCount {
		return "1"
	} else {
		return "0"
	}
}
func leastCommonBit(input []string, index int) string {
	oneCount := 0
	zeroCount := 0
	for _, line := range input {
		if line[index] == '1' {
			oneCount++
		} else {
			zeroCount++
		}
	}
	if oneCount == zeroCount {
		return "0"
	}
	if oneCount > zeroCount {
		return "0"
	} else {
		return "1"
	}
}
