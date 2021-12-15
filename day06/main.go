package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// 386755
	fmt.Printf("Part 1 result: %d\n", fishesCount("day06/input.txt", 80))
	// 18065
	fmt.Printf("Part 2 result: %d\n", fishesCountM("day06/input.txt", 256))
}

// Iterative
func fishesCount(path string, days int) int {
	fishes := parse(path)
	return len(getFishes(fishes, days))
}

func getFishes(fishes []int, days int) []int {
	result := append([]int{}, fishes...)
	for i := 0; i < days; i++ {
		for j, fish := range result {
			if fish != 0 {
				result[j] = fish - 1
				continue
			}
			result[j] = 6
			result = append(result, 8)
		}

	}

	return result
}

// Iterative with map
func fishesCountM(path string, days int) int {
	fishes := parse(path)
	return getFishesCountM(fishes, days)
}

func getFishesCountM(fishes []int, days int) int {
	fishesByTimer := getFishesMap(fishes, days)
	var totalFishes int
	for fish, count := range fishesByTimer {
		// don't totalFishes the staged fishes
		if fish == -1 {
			continue
		}
		totalFishes += count
	}

	return totalFishes
}

func getFishesMap(fishes []int, days int) map[int]int {
	// store fishes by value (from -1 to 8)
	// -1 is used as step to rotate the value
	result := map[int]int{-1: 0, 0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

	// initialize fishes map
	for _, fish := range fishes {
		result[fish] = result[fish] + 1
	}

	for i := 0; i < days; i++ {
		// decrease the timer
		for i := 0; i <= 9; i++ {
			result[i-1] = result[i]
		}
		toUp := result[-1]
		result[8] = toUp
		result[6] = result[6] + toUp
	}

	return result
}

// Recursive
func countFishesRec(fishes []int, days int) int {
	var count int
	for _, fish := range fishes {
		count += countFish(fish, days)
	}

	return count
}

func countFish(timer, last int) int {
	days := last - timer
	if days <= 0 {
		return 1
	}
	// as the timer is at 0, we can
	// * create a new fish with 8 days to create a new fish
	// * continue the fish life with 6 days (the next day) before creating new fish.
	return countFish(6, days-1) + countFish(8, days-1)
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
