package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	positions := New("test.txt")
	//assert.Equal(t, 37, getCheapestFuelOutcome(positions))
	assert.Equal(t, 37, positions.CheapestFuelCost(Abs))
}

func TestPart2(t *testing.T) {
	positions := New("test.txt")
	assert.Equal(t, 168, positions.CheapestFuelCost(Triangle))
}

func TestCumulCost(t *testing.T) {
	tests := []struct {
		desc     string
		from     int
		to       int
		expected int
	}{
		{
			desc:     "Move from 4 to 4: 1 fuel",
			from:     4,
			to:       4,
			expected: 0,
		},
		{
			desc:     "Move from 4 to 5: 1 fuel",
			from:     4,
			to:       5,
			expected: 1,
		},
		{
			desc:     "Move from 7 to 5: 3 fuel",
			from:     7,
			to:       5,
			expected: 3,
		},
		{
			desc:     "Move from 2 to 5: 6 fuel",
			from:     2,
			to:       5,
			expected: 6,
		},
		{
			desc:     "Move from 1 to 5: 10 fuel",
			from:     1,
			to:       5,
			expected: 10,
		},
		{
			desc:     "Move from 14 to 5: 45 fuel",
			from:     14,
			to:       5,
			expected: 45,
		},
		{
			desc:     "Move from 16 to 5: 66 fuel",
			from:     16,
			to:       5,
			expected: 66,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, Triangle(test.from-test.to))
		})
	}
}
