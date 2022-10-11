package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputArr := parseInput(input)
	println("Day One: ", simulateDay(inputArr, 0, 80))
	println("Day Two: ", simulateDay(inputArr, 0, 256))

}

// shoutout alexchao for this solution
func simulateDay(inputArr []int, day int, length int) int {
	fishies := make([]int, 9)

	for _, v := range inputArr {
		fishies[v]++
	}

	for i := 0; i < length; i++ {
		newBabbies := fishies[0]
		for j := 0; j < len(fishies)-1; j++ {
			fishies[j] = fishies[j+1]
		}
		fishies[8] = newBabbies
		fishies[6] += newBabbies
	}

	totalFishies := 0
	for _, v := range fishies {
		totalFishies += v
	}
	return totalFishies
}

func parseInput(input string) []int {
	temp := strings.Split(input, ",")
	values := make([]int, len(temp))
	for i, v := range temp {
		val, _ := strconv.Atoi(v)
		values[i] = val
	}
	return values
}
