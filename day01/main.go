package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	measures := getMeasures()

	fmt.Printf("Part 1 Result: %d\n", larger(measures))
	fmt.Printf("Part 2 Result: %d\n", sumLarger(measures))
}

func sumLarger(numbers []int) int {
	count := 0
	for i := 3; i < len(numbers); i++ {
		if numbers[i-3]+numbers[i-2]+numbers[i-1] < numbers[i-2]+numbers[i-1]+numbers[i] {
			count++
		}
	}

	return count
}

func larger(numbers []int) int {
	larger := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			larger++
		}
	}

	return larger
}

func getMeasures() []int {
	b, err := ioutil.ReadFile("day01/input.txt")
	if err != nil {
		fmt.Printf("cannot read file %v", err)
	}

	lines := strings.Split(string(b), "\n")
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			panic("Oops")
		}
		nums = append(nums, n)
	}

	return nums
}
