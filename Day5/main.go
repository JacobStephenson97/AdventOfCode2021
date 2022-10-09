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
	diagram := createDiagram(input)
	inputArr := strings.Split(input, "\r\n")

	for _, line := range inputArr {
		pointArr := parseLine(line)
		for _, point := range pointArr {
			diagram[point[1]][point[0]]++
		}
	}
	count := 0
	for _, line := range diagram {
		for _, point := range line {
			if point > 1 {
				count++
			}
		}
	}
	println(count)
}

func createDiagram(input string) [][]int {
	diagram := make([][]int, getMaxNumber(input)+1)
	for i := range diagram {
		diagram[i] = make([]int, getMaxNumber(input)+1)
	}
	return diagram
}

func parseLine(line string) [][]int {
	coords := strings.Split(line, " ")
	pointOne := coords[0]
	pointTwo := coords[2]

	x1, y1 := parsePoint(pointOne)
	x2, y2 := parsePoint(pointTwo)

	if x1 == x2 {
		// vertical line
		length := int(math.Abs(float64(y1) - float64(y2)))
		pointArr := make([][]int, length+1)

		for i, _ := range pointArr {
			if y1 <= y2 {
				pointArr[i] = []int{x1, y1 + i}
			} else {
				pointArr[i] = []int{x1, y1 - i}
			}
		}
		return pointArr
	} else if y1 == y2 {
		length := int(math.Abs(float64(x1) - float64(x2)))
		pointArr := make([][]int, length+1)
		for i, _ := range pointArr {
			if x1 <= x2 {
				pointArr[i] = []int{x1 + i, y1}
			} else {
				pointArr[i] = []int{x1 - i, y1}
			}
		}
		return pointArr
	}
	return nil
}

func parsePoint(point string) (int, int) {
	coords := strings.Split(point, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return x, y
}
func getMaxNumber(input string) int {
	max := 0
	inputArr := strings.Split(input, "\r\n")
	for _, line := range inputArr {
		coords := strings.Split(line, " ")
		pointOne := coords[0]
		pointTwo := coords[2]
		for _, point := range []string{pointOne, pointTwo} {
			x, y := parsePoint(point)
			if x > max {
				max = x
			}
			if y > max {
				max = y
			}
		}
	}
	return max
}
