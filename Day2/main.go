package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	day2(input)
}

func day2(input string) {
	maneuvers := strings.Split(input, "\r")

	x, y, aim := 0, 0, 0

	for _, maneuver := range maneuvers {
		dir, dist := parseInput(maneuver)
		dir = strings.TrimSpace(dir)

		switch dir {
		case "forward":
			x += dist
			if aim != 0 {
				y += dist * aim
			}
		case "down":
			aim += dist
		case "up":
			aim -= dist
		}
	}
	println(x * y)
}

func parseInput(input string) (string, int) {
	var dir = strings.Split(input, " ")[0]
	var dist, err = strconv.Atoi(strings.Split(input, " ")[1])
	if err != nil {
		panic(err)
	}
	return dir, dist
}
