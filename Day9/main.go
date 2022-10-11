package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	sum := 0
	parsedInput := parseInput(input)
	lowPoints := [][]int{}
	for i, line := range parsedInput {
		for j := range line {
			if checkAdjecent(parsedInput, i, j) {
				num := parsedInput[i][j]
				lowPoints = append(lowPoints, []int{i, j})
				sum += num + 1
			}
		}
	}
	biggest := make([]int, 3)
	for _, lowPoint := range lowPoints {
		basin := [][]int{lowPoint}
		basin = addToBasin(lowPoint, basin, parsedInput)
		min, index := getMin(biggest)
		if len(basin) > min {
			biggest[index] = len(basin)
		}
	}
	println(biggest[0] * biggest[1] * biggest[2])
}

func addToBasin(lowPoint []int, basin [][]int, parsedInput [][]int) [][]int {
	x := lowPoint[0]
	y := lowPoint[1]

	if x > 0 {
		if parsedInput[x-1][y] > parsedInput[x][y] && parsedInput[x-1][y] != 9 {
			if !contains(basin, []int{x - 1, y}) {
				basin = append(basin, []int{x - 1, y})
				basin = addToBasin([]int{x - 1, y}, basin, parsedInput)
			}
		}
	}
	if x < len(parsedInput)-1 {
		if parsedInput[x+1][y] > parsedInput[x][y] && parsedInput[x+1][y] != 9 {
			if !contains(basin, []int{x + 1, y}) {
				basin = append(basin, []int{x + 1, y})
				basin = addToBasin([]int{x + 1, y}, basin, parsedInput)
			}
		}
	}
	if y > 0 {
		if parsedInput[x][y-1] > parsedInput[x][y] && parsedInput[x][y-1] != 9 {
			if !contains(basin, []int{x, y - 1}) {
				basin = append(basin, []int{x, y - 1})
				basin = addToBasin([]int{x, y - 1}, basin, parsedInput)
			}
		}
	}
	if y < len(parsedInput[x])-1 {
		if parsedInput[x][y+1] > parsedInput[x][y] && parsedInput[x][y+1] != 9 {
			if !contains(basin, []int{x, y + 1}) {
				basin = append(basin, []int{x, y + 1})
				basin = addToBasin([]int{x, y + 1}, basin, parsedInput)
			}
		}
	}
	return basin
}

func parseInput(input string) [][]int {
	inputSlice := strings.Split(input, "\r\n")
	parsedInput := make([][]int, len(inputSlice))
	for i, line := range inputSlice {
		chars := strings.Split(line, "")
		temp := make([]int, len(chars))
		for j, char := range chars {
			temp[j], _ = strconv.Atoi(char)
		}
		parsedInput[i] = temp
	}
	return parsedInput
}
func checkAdjecent(parsedInput [][]int, x int, y int) bool {
	if x == 0 {
		if parsedInput[x+1][y] <= parsedInput[x][y] {
			return false
		}
		if y != len(parsedInput[x])-1 {
			if parsedInput[x][y+1] <= parsedInput[x][y] {
				return false
			}
		}
		if y != 0 {
			if parsedInput[x][y-1] <= parsedInput[x][y] {
				return false
			}
		}
	} else if x == len(parsedInput)-1 {
		if parsedInput[x-1][y] <= parsedInput[x][y] {
			return false
		}
		if y != len(parsedInput[x])-1 {
			if parsedInput[x][y+1] <= parsedInput[x][y] {
				return false
			}
		}
		if y != 0 {
			if parsedInput[x][y-1] <= parsedInput[x][y] {
				return false
			}
		}
	} else {
		if parsedInput[x+1][y] <= parsedInput[x][y] {
			return false
		}
		if parsedInput[x-1][y] <= parsedInput[x][y] {
			return false
		}
		if y != len(parsedInput[x])-1 {
			if parsedInput[x][y+1] <= parsedInput[x][y] {
				return false
			}
		}
		if y != 0 {
			if parsedInput[x][y-1] <= parsedInput[x][y] {
				return false
			}
		}
	}
	return true
}
func contains(s [][]int, e []int) bool {
	for _, a := range s {
		if a[0] == e[0] && a[1] == e[1] {
			return true
		}
	}
	return false
}
func getMin(biggest []int) (int, int) {
	min := biggest[0]
	lowestIndex := 0
	for i, num := range biggest {
		if num < min {
			min = num
			lowestIndex = i
		}
	}
	return min, lowestIndex

}
