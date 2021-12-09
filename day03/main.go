package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Elements [][]int

func (e Elements) MostCommonBit() string {
	var mcb string

	for _, bits := range e {
		if len(bits) == 0 {
			continue
		}

		if bits[0] > bits[1] {
			mcb += "0"
		} else {
			mcb += "1"
		}
	}

	return mcb
}

func (e Elements) Gamma() int {
	gamma, err := strconv.ParseInt(e.MostCommonBit(), 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(gamma)
}

func (e Elements) LessCommonBit() string {
	var lcb string

	for _, bits := range e {
		if len(bits) == 0 {
			continue
		}

		if bits[0] <= bits[1] {
			lcb += "0"
		} else {
			lcb += "1"
		}
	}

	return lcb
}

func (e Elements) Epsilon() int {
	epsilon, err := strconv.ParseInt(e.LessCommonBit(), 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(epsilon)
}

func (e Elements) ResultPart1() int {
	return e.Gamma() * e.Epsilon()
}

func filterNumbers(numbers []string, prefix string) []string {
	var filtered []string
	for _, number := range numbers {
		if strings.HasPrefix(number, prefix) {
			filtered = append(filtered, number)
		}
	}

	return filtered
}

func Oxygen(numbers []string) int {
	var mcb string
	filteredNumbers := append([]string{}, numbers...)

	for i := 0; len(filteredNumbers) != 1; i++ {
		oxygenElts := parse(filteredNumbers)
		mcb += string(oxygenElts.MostCommonBit()[i])
		filteredNumbers = filterNumbers(filteredNumbers, mcb)
	}

	if len(filteredNumbers) != 1 {
		panic("Expected 1 as filteredNumbers len")
	}

	return toInt(filteredNumbers[0])
}

func CO2(numbers []string) int {
	var lcb string
	filteredNumbers := append([]string{}, numbers...)

	for i := 0; len(filteredNumbers) != 1; i++ {
		co2Elts := parse(filteredNumbers)
		lcb += string(co2Elts.LessCommonBit()[i])
		filteredNumbers = filterNumbers(filteredNumbers, lcb)
	}

	if len(filteredNumbers) != 1 {
		panic("Expected 1 as filteredNumbers len")
	}

	return toInt(filteredNumbers[0])
}

// 6677951
func ResultPart2(numbers []string) int {
	return Oxygen(numbers) * CO2(numbers)
}

// Gamma * Epsilon rate: 2743844
func main() {
	numbers := getNumbers("day03/input.txt")
	elements := parse(numbers)
	fmt.Printf("Gamma * Epsilon rate: %d\n", elements.ResultPart1())
	fmt.Printf("Oxygen * CO2 rate: %d\n", ResultPart2(numbers))
}

func parse(numbers []string) Elements {
	// 0111 1110 1011
	// [12][2]
	elements := make(Elements, len(numbers[0]))
	for _, n := range numbers {
		for i := 0; i < len(n); i++ {
			if len(elements[i]) == 0 {
				elements[i] = []int{0, 0}
			}

			val := int(n[i] - '0')
			elements[i][val] = elements[i][val] + 1
		}
	}

	return elements
}

func getNumbers(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("cannot read file %v", err)
	}

	lines := strings.Split(string(b), "\n")
	moves := make([]string, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		moves = append(moves, l)
	}

	return moves
}

func toInt(bits string) int {
	value, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(value)
}
