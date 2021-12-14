package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 5934, resultPart1("test.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 12, resultPart2("test.txt"))
}

// -- && 0 -> 6 && add 8

func TestParse(t *testing.T) {
	fishes := parse("test.txt")
	assert.Equal(t, []int{3, 4, 3, 1, 2}, fishes)
}
