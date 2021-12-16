package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	positions := New("day07/input.txt")

	fmt.Printf("Part 1 result: %d\n", positions.CheapestFuelCost(Abs))
	fmt.Printf("Part 2 result: %d\n", positions.CheapestFuelCost(Triangle))
}

type Positions []int

func New(path string) Positions {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("cannot read file %v", err)
	}

	var positions []int
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		numbers := strings.Split(line, ",")
		for _, nb := range numbers {
			positions = append(positions, toInt(nb))
		}
	}

	return positions
}

type CostComputer func(distance int) int

func (p Positions) CheapestFuelCost(getCost CostComputer) int {
	sort.Ints(p)
	min := p[0]
	max := p[len(p)-1]

	var costs []int
	for to := min; to <= max; to++ {
		var cost int
		for _, from := range p {
			cost += getCost(from - to)
		}
		costs = append(costs, cost)
	}

	sort.Ints(costs)
	return costs[0]
}

func Abs(n int) int {
	if n >= 0 {
		return n
	}

	return -n
}

func Triangle(distance int) int {
	n := Abs(distance)
	return (n * (n + 1)) / 2
}

func toInt(s string) int {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		panic("Boom")
	}

	return int(value)
}
