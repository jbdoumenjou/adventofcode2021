package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	positions := parse("day07/input.txt")

	fmt.Printf("Part 1 result: %d\n", getcheapestFuelOutcome(positions))
}

func getcheapestFuelOutcome(positions []int) int {
	var outcomes []int
	for _, p := range positions {
		var cost int
		for _, position := range positions {
			if p == position {
				continue
			}
			// abs
			if position > p {
				cost += position - p
				continue
			}
			cost += p - position
		}
		outcomes = append(outcomes, cost)
	}

	sort.Ints(outcomes)

	return outcomes[0]
}

// Tools
func parse(path string) []int {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("cannot read file %v", err)
	}

	var fishes []int
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		numbers := strings.Split(line, ",")
		for _, nb := range numbers {
			fishes = append(fishes, toInt(nb))
		}
	}

	return fishes
}

func toInt(s string) int {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		panic("Boom")
	}

	return int(value)
}
