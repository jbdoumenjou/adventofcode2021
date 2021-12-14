package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func resultPart1(path string) int {
	return 0
}

func resultPart2(path string) int {
	return 0
}

func main() {
	// 6461
	fmt.Printf("Part 1 result: %d\n", resultPart1("day06/input.txt"))
	// 18065
	fmt.Printf("Part 2 result: %d\n", resultPart2("day06/input.txt"))
}

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
