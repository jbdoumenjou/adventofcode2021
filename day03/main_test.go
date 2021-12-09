package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestPart1(t *testing.T) {
	numbers := getNumbers("test.txt")
	elements := parse(numbers)
	assert.Equal(t, elements.ResultPart1(), 198)
}

func TestPart2(t *testing.T) {
	numbers := getNumbers("test.txt")
	assert.Equal(t, ResultPart2(numbers), 230)
}
