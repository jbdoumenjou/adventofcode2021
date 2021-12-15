package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	//assert.Equal(t, 5934, fishesCount("test.txt", 80))
	assert.Equal(t, 5934, fishesCountM("test.txt", 80))
}

func TestPart2(t *testing.T) {
	//assert.Equal(t, 26984457539, fishesCount("test.txt", 256))
	assert.Equal(t, 26984457539, fishesCountM("test.txt", 256))
}

func TestProcess(t *testing.T) {
	expectedFishesByDay := [][]int{
		// 0
		{3, 4, 3, 1, 2},
		// 1
		{2, 3, 2, 0, 1},
		// 2
		{1, 2, 1, 6, 0, 8},
		// 3
		{0, 1, 0, 5, 6, 7, 8},
		// 4
		{6, 0, 6, 4, 5, 6, 7, 8, 8},
		// 5
		{5, 6, 5, 3, 4, 5, 6, 7, 7, 8},
		// 6
		{4, 5, 4, 2, 3, 4, 5, 6, 6, 7},
		// 7
		{3, 4, 3, 1, 2, 3, 4, 5, 5, 6},
		// 8
		{2, 3, 2, 0, 1, 2, 3, 4, 4, 5},
		// 9
		{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8},
		// 10
		{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 7, 8},
		// 11
		{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 7, 8, 8, 8},
		// 12
		{5, 6, 5, 3, 4, 5, 6, 0, 0, 1, 5, 6, 7, 7, 7, 8, 8},
		// 13
		{4, 5, 4, 2, 3, 4, 5, 6, 6, 0, 4, 5, 6, 6, 6, 7, 7, 8, 8},
		// 14
		{3, 4, 3, 1, 2, 3, 4, 5, 5, 6, 3, 4, 5, 5, 5, 6, 6, 7, 7, 8},
		// 15
		{2, 3, 2, 0, 1, 2, 3, 4, 4, 5, 2, 3, 4, 4, 4, 5, 5, 6, 6, 7},
		// 16
		{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 1, 2, 3, 3, 3, 4, 4, 5, 5, 6, 8},
		// 17
		{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 0, 1, 2, 2, 2, 3, 3, 4, 4, 5, 7, 8},
		// 18
		{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8},
	}

	for day, fishes := range expectedFishesByDay {
		assert.Equal(t, fishes, getFishes(expectedFishesByDay[0], day))
	}
}

func TestCountFishes(t *testing.T) {
	expectedFishesByDay := [][]int{
		// 0
		{3, 4, 3, 1, 2},
		// 1
		{2, 3, 2, 0, 1},
		// 2
		{1, 2, 1, 6, 0, 8},
		// 3
		{0, 1, 0, 5, 6, 7, 8},
		// 4
		{6, 0, 6, 4, 5, 6, 7, 8, 8},
		// 5
		{5, 6, 5, 3, 4, 5, 6, 7, 7, 8},
		// 6
		{4, 5, 4, 2, 3, 4, 5, 6, 6, 7},
		// 7
		{3, 4, 3, 1, 2, 3, 4, 5, 5, 6},
		// 8
		{2, 3, 2, 0, 1, 2, 3, 4, 4, 5},
		// 9
		{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8},
		// 10
		{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 7, 8},
		// 11
		{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 7, 8, 8, 8},
		// 12
		{5, 6, 5, 3, 4, 5, 6, 0, 0, 1, 5, 6, 7, 7, 7, 8, 8},
		// 13
		{4, 5, 4, 2, 3, 4, 5, 6, 6, 0, 4, 5, 6, 6, 6, 7, 7, 8, 8},
		// 14
		{3, 4, 3, 1, 2, 3, 4, 5, 5, 6, 3, 4, 5, 5, 5, 6, 6, 7, 7, 8},
		// 15
		{2, 3, 2, 0, 1, 2, 3, 4, 4, 5, 2, 3, 4, 4, 4, 5, 5, 6, 6, 7},
		// 16
		{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 1, 2, 3, 3, 3, 4, 4, 5, 5, 6, 8},
		// 17
		{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 0, 1, 2, 2, 2, 3, 3, 4, 4, 5, 7, 8},
		// 18
		{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8},
	}

	for day, fishes := range expectedFishesByDay {
		//assert.Equal(t, fishes, getFishes(expectedFishesByDay[0], day))
		//assert.Equal(t, len(fishes), countFishesRec(expectedFishesByDay[0], day))
		assert.Equal(t, len(fishes), getFishesCountM(expectedFishesByDay[0], day))
	}
}

func TestParse(t *testing.T) {
	fishes := parse("test.txt")
	assert.Equal(t, []int{3, 4, 3, 1, 2}, fishes)
}

func TestByFish(t *testing.T) {
	tests := []struct {
		desc     string
		start    int
		last     int
		expected int
	}{
		{
			desc:     "fish of 3 after the first day",
			start:    3,
			last:     1,
			expected: 1,
		},
		{
			desc:     "fish of 3 day 3",
			start:    3,
			last:     3,
			expected: 1,
		},
		{
			desc:     "fish of 3 day 4",
			start:    3,
			last:     4,
			expected: 2,
		},
		{
			desc:     "fish of 3 day 10",
			start:    3,
			last:     10,
			expected: 2,
		},
		{
			desc:     "fish of 3 day 11",
			start:    3,
			last:     11,
			expected: 3,
		},
		{
			desc:     "fish of 3 day 12",
			start:    3,
			last:     12,
			expected: 3,
		},
		{
			desc:     "fish of 3 day 13",
			start:    3,
			last:     13,
			expected: 4,
		},
		{
			desc:     "fish of 3 day 17",
			start:    3,
			last:     17,
			expected: 4,
		},
		{
			desc:     "fish of 3 day 18",
			start:    3,
			last:     18,
			expected: 5,
		},
	}

	expectedFishesByDay := [][]int{
		{3},    // 0
		{2},    // 1
		{1},    // 2
		{0},    // 3
		{6, 8}, // 4
		{5, 7},
		{4, 6},
		{3, 5},
		{2, 4},
		{1, 3},       // 9
		{0, 2},       // 10
		{6, 1, 8},    // 11
		{5, 0, 7},    // 12
		{4, 6, 6, 8}, // 13
		{3, 5, 5, 7},
		{2, 4, 4, 6},
		{1, 3, 3, 5},    // 16
		{0, 2, 2, 4},    // 17
		{6, 1, 1, 3, 8}, // 18
	}
	_ = expectedFishesByDay

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, countFish(test.start, test.last))
		})
	}
}

func TestGetFishesM(t *testing.T) {
	tests := []struct {
		desc     string
		fishes   []int
		days     int
		expected map[int]int
	}{
		{
			desc:     "one fish of 3 after 1 day",
			fishes:   []int{3},
			days:     1,
			expected: map[int]int{-1: 0, 0: 0, 1: 0, 2: 1, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		},
		{
			desc:     "one fish of 3 after 3 days",
			fishes:   []int{3},
			days:     3,
			expected: map[int]int{-1: 0, 0: 1, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		},
		{
			desc:     "one fish of 3 after 4 days",
			fishes:   []int{3},
			days:     4,
			expected: map[int]int{-1: 1, 0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 1, 7: 0, 8: 1},
		},
		{
			desc:     "one fish of 3 after 5 days",
			fishes:   []int{3},
			days:     5,
			expected: map[int]int{-1: 0, 0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 1, 6: 0, 7: 1, 8: 0},
		},
		{
			desc:   "some fishes after 0 day",
			fishes: []int{3, 4, 3, 1, 2},
			days:   0,
			//{3, 4, 3, 1, 2},
			expected: map[int]int{-1: 0, 0: 0, 1: 1, 2: 1, 3: 2, 4: 1, 5: 0, 6: 0, 7: 0, 8: 0},
		},
		{
			desc:   "tests fishes after 1 day",
			fishes: []int{3, 4, 3, 1, 2},
			days:   1,
			// {2, 3, 2, 0, 1},
			expected: map[int]int{-1: 0, 0: 1, 1: 1, 2: 2, 3: 1, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		},
		{
			desc:   "tests fishes after 2 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   2,
			// {1, 2, 1, 6, 0, 8},
			expected: map[int]int{-1: 1, 0: 1, 1: 2, 2: 1, 3: 0, 4: 0, 5: 0, 6: 1, 7: 0, 8: 1},
		},
		{
			desc:   "tests fishes after 3 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   3,
			// {0, 1, 0, 5, 6, 7, 8}
			expected: map[int]int{-1: 1, 0: 2, 1: 1, 2: 0, 3: 0, 4: 0, 5: 1, 6: 1, 7: 1, 8: 1},
		},
		{
			desc:   "tests fishes after 4 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   4,
			// {6, 0, 6, 4, 5, 6, 7, 8, 8}
			expected: map[int]int{-1: 2, 0: 1, 1: 0, 2: 0, 3: 0, 4: 1, 5: 1, 6: 3, 7: 1, 8: 2},
		},
		{
			desc:   "tests fishes after 5 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   5,
			// {5, 6, 5, 3, 4, 5, 6, 7, 7, 8}
			expected: map[int]int{-1: 1, 0: 0, 1: 0, 2: 0, 3: 1, 4: 1, 5: 3, 6: 2, 7: 2, 8: 1},
		},
		{
			desc:   "tests fishes after 6 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   6,
			// {4, 5, 4, 2, 3, 4, 5, 6, 6, 7}
			expected: map[int]int{-1: 0, 0: 0, 1: 0, 2: 1, 3: 1, 4: 3, 5: 2, 6: 2, 7: 1, 8: 0},
		},
		{
			desc:   "tests fishes after 7 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   7,
			// {3, 4, 3, 1, 2, 3, 4, 5, 5, 6}
			expected: map[int]int{-1: 0, 0: 0, 1: 1, 2: 1, 3: 3, 4: 2, 5: 2, 6: 1, 7: 0, 8: 0},
		},
		{
			desc:   "tests fishes after 8 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   8,
			// {2, 3, 2, 0, 1, 2, 3, 4, 4, 5}
			expected: map[int]int{-1: 0, 0: 1, 1: 1, 2: 3, 3: 2, 4: 2, 5: 1, 6: 0, 7: 0, 8: 0},
		},
		{
			desc:   "tests fishes after 9 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   9,
			// {1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8}
			expected: map[int]int{-1: 1, 0: 1, 1: 3, 2: 2, 3: 2, 4: 1, 5: 0, 6: 1, 7: 0, 8: 1},
		},
		{
			desc:   "tests fishes after 10 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   10,
			// {0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 7, 8}
			expected: map[int]int{-1: 1, 0: 3, 1: 2, 2: 2, 3: 1, 4: 0, 5: 1, 6: 1, 7: 1, 8: 1},
		},
		{
			desc:   "tests fishes after 11 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   11,
			// {6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 7, 8, 8, 8}
			expected: map[int]int{-1: 3, 0: 2, 1: 2, 2: 1, 3: 0, 4: 1, 5: 1, 6: 4, 7: 1, 8: 3},
		},
		{
			desc:   "tests fishes after 12 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   12,
			// {5, 6, 5, 3, 4, 5, 6, 0, 0, 1, 5, 6, 7, 7, 7, 8, 8}
			expected: map[int]int{-1: 2, 0: 2, 1: 1, 2: 0, 3: 1, 4: 1, 5: 4, 6: 3, 7: 3, 8: 2},
		},
		{
			desc:   "tests fishes after 13 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   13,
			// {4, 5, 4, 2, 3, 4, 5, 6, 6, 0, 4, 5, 6, 6, 6, 7, 7, 8, 8}
			expected: map[int]int{-1: 2, 0: 1, 1: 0, 2: 1, 3: 1, 4: 4, 5: 3, 6: 5, 7: 2, 8: 2},
		},
		{
			desc:   "tests fishes after 14 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   14,
			// {3, 4, 3, 1, 2, 3, 4, 5, 5, 6, 3, 4, 5, 5, 5, 6, 6, 7, 7, 8}
			expected: map[int]int{-1: 1, 0: 0, 1: 1, 2: 1, 3: 4, 4: 3, 5: 5, 6: 3, 7: 2, 8: 1},
		},
		{
			desc:   "tests fishes after 15 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   15,
			// {2, 3, 2, 0, 1, 2, 3, 4, 4, 5, 2, 3, 4, 4, 4, 5, 5, 6, 6, 7}
			expected: map[int]int{-1: 0, 0: 1, 1: 1, 2: 4, 3: 3, 4: 5, 5: 3, 6: 2, 7: 1, 8: 0},
		},
		{
			desc:   "tests fishes after 16 days",
			fishes: []int{3, 4, 3, 1, 2},
			days:   16,
			// {1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 1, 2, 3, 3, 3, 4, 4, 5, 5, 6, 8}
			expected: map[int]int{-1: 1, 0: 1, 1: 4, 2: 3, 3: 5, 4: 3, 5: 2, 6: 2, 7: 0, 8: 1},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, getFishesMap(test.fishes, test.days))
		})
	}
}
