package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	drawingNumbers, boards := parseInput(input)

	var bingoBoards = make([][][]int, len(boards))
	for i := range bingoBoards {
		bingoBoards[i] = make([][]int, len(boards[i]))
		for j := range bingoBoards[i] {
			bingoBoards[i][j] = make([]int, len(boards[i][j]))
		}
	}
	loopBoards(boards, drawingNumbers, bingoBoards)
}

func parseInput(input string) ([]string, [][][]string) {
	drawingNumbers := strings.Split(input, "\r")[0]

	allBoards := strings.Split(input, "\r\n")[1:]
	println(allBoards[0])
	numberSlice := strings.Split(drawingNumbers, ",")

	boardSlice := createBoards(allBoards)

	return numberSlice, boardSlice
}
func createBoards(allBoards []string) [][][]string {
	boardCount := -1
	var boards [][][]string
	var board [][]string
	for _, line := range allBoards {
		if line == "" {
			if boardCount != -1 {
				boards = append(boards, board)
				board = nil
			}
			boardCount++
		} else {
			if line != "" {
				numbers := strings.Split(line, " ")
				var lineNumbers []string
				for _, number := range numbers {
					if number != "" {
						lineNumbers = append(lineNumbers, number)
					}
				}
				board = append(board, lineNumbers)
			}
		}
	}
	return boards
}
func checkBingo(board [][][]int) (int, int) {
	for i := range board {
	row:
		for j := range board[i] {
			for k := range board[i][j] {
				if board[i][j][k] == 0 {
					continue row
				}
			}
			return i, j
		}
	}
	return 0, 0
}
func checkBingoY(board [][][]int) (int, int) {
	for i := range board {
	row:
		for j := range board[i] {
			for k := range board[i][j] {
				if board[i][k][j] == 0 {
					continue row
				}
			}
			return i, j
		}
	}
	return 0, 0
}
func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
func remove(s [][][]string, i int) [][][]string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func removeInt(s [][][]int, i int) [][][]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func loopBoards(boards [][][]string, drawingNumbers []string, bingoBoards [][][]int) int {
	for index, number := range drawingNumbers {
		for i := range boards {
			for j := range boards[i] {
				for k := range boards[i][j] {
					if boards[i][j][k] == number {
						bingoBoards[i][j][k] = index + 1
						x, y := checkBingoY(bingoBoards)
						if x == 0 && y == 0 {
							x, y = checkBingo(bingoBoards)
						}
						if x != 0 || y != 0 {

							if len(boards) > 1 {
								boards = remove(boards, x)
								bingoBoards = removeInt(bingoBoards, x)
								return loopBoards(boards, drawingNumbers, bingoBoards)
							}
							var sum int
							for index, line := range boards[x] {
								for nindex, num := range line {
									if bingoBoards[x][index][nindex] == 0 {
										sum += parseInt(num)
									}
								}
							}
							println(sum * parseInt(number))
							return sum * parseInt(boards[x][y][k])
						}
					}
				}
			}
		}
	}
	return 0
}
