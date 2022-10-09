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
			diagram[point[0]][point[1]]++
		}
	}
	println(diagram)
}

func createDiagram(input string) [][]int {
	diagram := make([][]int, 10)
	for i := range diagram {
		diagram[i] = make([]int, 10)
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
