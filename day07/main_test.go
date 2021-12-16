package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	positions := parse("test.txt")
	assert.Equal(t, 37, getcheapestFuelOutcome(positions))
}

func TestPart2(t *testing.T) {
}
