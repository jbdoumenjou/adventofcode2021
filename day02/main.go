package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Move struct {
	Direction string
	Value     int
}

type Position struct {
	X   int
	Y   int
	Aim int
}

func (p *Position) applySimple(move Move) *Position {
	switch move.Direction {
	case "forward":
		p.X += move.Value
	case "down":
		p.Y += move.Value
	case "up":
		p.Y -= move.Value
	}

	return p
}

func (p *Position) applyWithAim(move Move) *Position {
	switch move.Direction {
	case "forward":
		p.X += move.Value
		p.Y += p.Aim * move.Value
	case "down":
		p.Aim += move.Value
	case "up":
		p.Aim -= move.Value
	}

	return p
}

func (p Position) Result() int {
	return p.X * p.Y
}

func main() {
	moves := getMoves()
	p := Position{}
	for _, move := range moves {
		p.applySimple(move)
	}

	fmt.Printf("Part 1 result: %d\n", p.Result())

	p = Position{}
	for _, move := range moves {
		p.applyWithAim(move)
	}

	fmt.Printf("Part 2 result: %d\n", p.Result())
}

func getMoves() []Move {
	b, err := ioutil.ReadFile("day02/input.txt")
	if err != nil {
		fmt.Printf("cannot read file %v", err)
	}

	lines := strings.Split(string(b), "\n")
	moves := make([]Move, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		rawMove := strings.Split(l, " ")
		m := Move{
			Direction: rawMove[0],
		}

		n, err := strconv.Atoi(rawMove[1])
		if err != nil {
			panic("Oops")
		}
		m.Value = n

		moves = append(moves, m)
	}

	return moves
}
