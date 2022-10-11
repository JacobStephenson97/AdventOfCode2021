package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputArr := parseInput(input)

	var min int
	println(calcFuelPartTwo(inputArr, 5))
	for i := 0; i < getMax(inputArr); i++ {
		fuel := calcFuelPartTwo(inputArr, i)
		if min == 0 || fuel < min {
			min = fuel
		}
	}
	println(min)
}

func calcFuelDayOne(inputArr []int, target int) int {
	sum := 0
	for _, v := range inputArr {
		sum += int(math.Abs(float64(v) - float64(target)))
	}
	return sum
}

func calcFuelPartTwo(inputArr []int, target int) int {
	sum := 0
	for _, v := range inputArr {
		for i := 0; i < int(math.Abs(float64(v)-float64(target))); i++ {
			sum += 1 * (i + 1)
		}
	}
	return sum
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
func getMax(input []int) int {
	max := 0
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}
