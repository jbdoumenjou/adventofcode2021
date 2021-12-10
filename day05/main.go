package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func resultPart1(path string) int {
	points := parse(path, false)
	var count int
	for _, point := range points {
		if point >= 2 {
			count++
		}
	}

	return count
}

func resultPart2(path string) int {
	points := parse(path, true)
	var count int
	for _, point := range points {
		if point >= 2 {
			count++
		}
	}

	return count
}

func main() {
	// 6461
	fmt.Printf("Part 1 result: %d\n", resultPart1("day05/input.txt"))
	// 18065
	fmt.Printf("Part 2 result: %d\n", resultPart2("day05/input.txt"))
}

type position struct {
	x int
	y int
}

// store the number of point for a given position
type posOccurrence map[position]int

func (p posOccurrence) print(size int) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			count := p[position{x: x, y: y}]
			if count == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", count)
			}
		}
		fmt.Println()

	}
}
func (p posOccurrence) AddSegment(segment string, allowDiagonal bool) {
	coords := strings.Split(segment, " -> ")
	if len(coords) != 2 {
		return
	}

	start := strings.Split(coords[0], ",")
	if len(start) != 2 {
		return
	}
	startPos := position{
		x: toInt(start[0]),
		y: toInt(start[1]),
	}

	end := strings.Split(coords[1], ",")
	if len(end) != 2 {
		return
	}
	endPos := position{
		x: toInt(end[0]),
		y: toInt(end[1]),
	}
	// For now, only consider horizontal and vertical lines:
	// lines where either x1 = x2 or y1 = y2.
	if startPos.x != endPos.x && startPos.y != endPos.y {
		if allowDiagonal {
			p.AddDiagonal(startPos, endPos)
		}
		return
	}

	beginX := startPos.x
	if startPos.x > endPos.x {
		beginX = endPos.x
	}
	endX := endPos.x
	if startPos.x > endPos.x {
		endX = startPos.x
	}

	beginY := startPos.y
	if startPos.y > endPos.y {
		beginY = endPos.y
	}
	endY := endPos.y
	if startPos.y > endPos.y {
		endY = startPos.y
	}

	for x := beginX; x <= endX; x++ {
		for y := beginY; y <= endY; y++ {
			count := p[position{x: x, y: y}]
			p[position{x: x, y: y}] = count + 1
		}
	}
}

func (p posOccurrence) AddDiagonal(begin, end position) {

	dirX := 1
	if begin.x > end.x {
		dirX = -1
	}
	dirY := 1
	if begin.y > end.y {
		dirY = -1
	}

	if begin.x != begin.y && begin.x != end.y &&
		dirX*(begin.x-end.x) != dirY*(begin.y-end.y) {
		return
	}

	x := begin.x
	y := begin.y
	for {
		count := p[position{x: x, y: y}]
		p[position{x: x, y: y}] = count + 1
		if x == end.x && y == end.y {
			break
		}
		x += dirX
		y += dirY
	}
}

func parse(path string, allowDiagonal bool) posOccurrence {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("cannot read file %v", err)
	}
	po := posOccurrence{}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		po.AddSegment(line, allowDiagonal)
	}

	return po
}

func toInt(s string) int {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		panic("Boom")
	}

	return int(value)
}
