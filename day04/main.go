package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func resultPart1(path string) int {
	numbers, boards := parse(path)
	for _, number := range numbers {
		for _, board := range boards {
			if board.Mark(number) {
				return board.SumUnmarked() * number
			}
		}
	}

	return 0
}

// not 15510
func main() {
	fmt.Printf("Part 1 result: %d\n", resultPart1("day04/input.txt"))
}

type Box struct {
	Nb     int
	Marked bool
}

type BoardRow []Box

// Mark the matching value and returns true if a whole line is marked
func (bl BoardRow) Mark(nb int) bool {
	bingo := true
	for i, box := range bl {
		if box.Nb == nb {
			box.Marked = true
			bl[i] = box
		}

		if !box.Marked {
			bingo = false
		}
	}

	return bingo
}

func (bl BoardRow) SumUnmarked() int {
	var sum int

	for _, box := range bl {
		if !box.Marked {
			sum += box.Nb
		}
	}

	return sum
}

type Board []BoardRow

// Mark the matching value and returns true if a whole line is marked
func (b Board) Mark(nb int) bool {
	// row
	for _, bl := range b {
		if bl.Mark(nb) {
			return true
		}
	}
	//col
	for rowIndex := 0; rowIndex < len(b[0]); rowIndex++ {
		bingo := true
		for colIndex := 0; colIndex < len(b); colIndex++ {
			if !b[colIndex][rowIndex].Marked {
				bingo = false
			}
		}
		if bingo {
			return true
		}
	}

	return false
}

func (b Board) SumUnmarked() int {
	var sum int
	for _, bl := range b {
		sum += bl.SumUnmarked()
	}

	return sum
}

func parseBoardLine(line string) []Box {
	var boardLine []Box

	v := strings.Split(line, " ")
	for _, nb := range v {
		if nb == "" {
			continue
		}

		boardLine = append(boardLine, Box{Nb: toInt(nb)})
	}

	return boardLine
}
func parseNumbersLine(line string) []int {
	var numbers []int

	v := strings.Split(line, ",")
	for _, nb := range v {
		if nb == "" {
			continue
		}

		numbers = append(numbers, toInt(nb))
	}

	return numbers
}

func parse(path string) ([]int, []Board) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("cannot read file %v", err)
	}

	var numbers []int
	lines := strings.Split(string(b), "\n")
	var boards []Board
	var currentBoard Board

	for i, line := range lines {
		if i == 0 {
			numbers = parseNumbersLine(line)
			continue
		}

		if line == "" {
			if currentBoard != nil && len(currentBoard) > 0 {
				boards = append(boards, currentBoard)
			}
			currentBoard = Board{}
			continue
		}

		currentBoard = append(currentBoard, parseBoardLine(line))
	}

	return numbers, boards
}

func toInt(bits string) int {
	value, err := strconv.ParseInt(bits, 10, 64)
	if err != nil {
		fmt.Println(err)
		panic("Boom")
	}

	return int(value)
}
