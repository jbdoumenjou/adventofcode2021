package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 5, resultPart1("test.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 12, resultPart2("test.txt"))
}

func TestPosOccurrence_AddSegment(t *testing.T) {
	tests := []struct {
		desc          string
		p             posOccurrence
		segment       string
		allowDiagonal bool
		expected      posOccurrence
	}{
		{
			desc:     "empty segment",
			p:        posOccurrence{},
			segment:  "",
			expected: posOccurrence{},
		},
		{
			desc:    "vertical positive segment",
			p:       posOccurrence{},
			segment: "0,9 -> 5,9",
			expected: posOccurrence{
				position{x: 0, y: 9}: 1,
				position{x: 1, y: 9}: 1,
				position{x: 2, y: 9}: 1,
				position{x: 3, y: 9}: 1,
				position{x: 4, y: 9}: 1,
				position{x: 5, y: 9}: 1},
		},
		{
			desc: "vertical positive segment with more than one point",
			p: posOccurrence{
				position{x: 1, y: 9}: 1,
			},
			segment: "0,9 -> 5,9",
			expected: posOccurrence{
				position{x: 0, y: 9}: 1,
				position{x: 1, y: 9}: 2,
				position{x: 2, y: 9}: 1,
				position{x: 3, y: 9}: 1,
				position{x: 4, y: 9}: 1,
				position{x: 5, y: 9}: 1},
		},
		{
			desc:     "Part1: ignore diagonals",
			p:        posOccurrence{},
			segment:  "0,9 -> 5,4",
			expected: posOccurrence{},
		},
		{
			desc:    "vertical negative segment",
			p:       posOccurrence{},
			segment: "5,9 -> 0,9",
			expected: posOccurrence{
				position{x: 0, y: 9}: 1,
				position{x: 1, y: 9}: 1,
				position{x: 2, y: 9}: 1,
				position{x: 3, y: 9}: 1,
				position{x: 4, y: 9}: 1,
				position{x: 5, y: 9}: 1},
		},
		{
			desc:    "horizontal positive segment",
			p:       posOccurrence{},
			segment: "0,1 -> 0,6",
			expected: posOccurrence{
				position{x: 0, y: 1}: 1,
				position{x: 0, y: 2}: 1,
				position{x: 0, y: 3}: 1,
				position{x: 0, y: 4}: 1,
				position{x: 0, y: 5}: 1,
				position{x: 0, y: 6}: 1},
		},
		{
			desc:    "horizontal negative segment",
			p:       posOccurrence{},
			segment: "0,6 -> 0,1",
			expected: posOccurrence{
				position{x: 0, y: 1}: 1,
				position{x: 0, y: 2}: 1,
				position{x: 0, y: 3}: 1,
				position{x: 0, y: 4}: 1,
				position{x: 0, y: 5}: 1,
				position{x: 0, y: 6}: 1},
		},
		{
			desc:          "diagonal positive segment",
			p:             posOccurrence{},
			segment:       "1,1 -> 4,4",
			allowDiagonal: true,
			expected: posOccurrence{
				position{x: 1, y: 1}: 1,
				position{x: 2, y: 2}: 1,
				position{x: 3, y: 3}: 1,
				position{x: 4, y: 4}: 1,
			},
		},
		{
			desc:          "diagonal negative segment",
			p:             posOccurrence{},
			segment:       "4,4 -> 1,1",
			allowDiagonal: true,
			expected: posOccurrence{
				position{x: 1, y: 1}: 1,
				position{x: 2, y: 2}: 1,
				position{x: 3, y: 3}: 1,
				position{x: 4, y: 4}: 1,
			},
		},
		{
			desc:          "diagonal negative segment",
			p:             posOccurrence{},
			segment:       "9,7 -> 7,9",
			allowDiagonal: true,
			expected: posOccurrence{
				position{x: 7, y: 9}: 1,
				position{x: 8, y: 8}: 1,
				position{x: 9, y: 7}: 1,
			},
		},
		{
			desc:          "diagonal negative segment",
			p:             posOccurrence{},
			segment:       "6,4 -> 2,0",
			allowDiagonal: true,
			expected: posOccurrence{
				position{x: 6, y: 4}: 1,
				position{x: 5, y: 3}: 1,
				position{x: 4, y: 2}: 1,
				position{x: 3, y: 1}: 1,
				position{x: 2, y: 0}: 1,
			},
		},
		{
			desc:          "diagonal negative segment",
			p:             posOccurrence{},
			segment:       "2,0 -> 6,4",
			allowDiagonal: true,
			expected: posOccurrence{
				position{x: 6, y: 4}: 1,
				position{x: 5, y: 3}: 1,
				position{x: 4, y: 2}: 1,
				position{x: 3, y: 1}: 1,
				position{x: 2, y: 0}: 1,
			},
		},
		{
			desc:          "diagonal negative segment",
			p:             posOccurrence{},
			segment:       "2,2 -> 0,4",
			allowDiagonal: true,
			expected: posOccurrence{
				position{x: 2, y: 2}: 1,
				position{x: 1, y: 3}: 1,
				position{x: 0, y: 4}: 1,
			},
		},
		{
			desc:          "diagonal negative segment",
			p:             posOccurrence{},
			segment:       "0,4 -> 2,2",
			allowDiagonal: true,
			expected: posOccurrence{
				position{x: 2, y: 2}: 1,
				position{x: 1, y: 3}: 1,
				position{x: 0, y: 4}: 1,
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			test.p.AddSegment(test.segment, test.allowDiagonal)
			assert.Len(t, test.p, len(test.expected))
			assert.Equal(t, test.expected, test.p)
		})
	}
}
