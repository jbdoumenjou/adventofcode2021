package main

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, resultPart1("test.txt"), 4512)
}

func TestParseBoardLine(t *testing.T) {
	line := parseBoardLine("22 13 17 11  0")
	assert.Equal(t, line, []int{22, 13, 17, 11, 0})
}

func TestParse(t *testing.T) {
	numbers, boards := parse("test.txt")
	assert.Equal(t, numbers, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1})
	assert.Equal(t, boards, []Board{
		{
			[]Box{{Nb: 22}, {Nb: 13}, {Nb: 17}, {Nb: 11}, {Nb: 0}},
			[]Box{{Nb: 8}, {Nb: 2}, {Nb: 23}, {Nb: 4}, {Nb: 24}},
			[]Box{{Nb: 21}, {Nb: 9}, {Nb: 14}, {Nb: 16}, {Nb: 7}},
			[]Box{{Nb: 6}, {Nb: 10}, {Nb: 3}, {Nb: 18}, {Nb: 5}},
			[]Box{{Nb: 1}, {Nb: 12}, {Nb: 20}, {Nb: 15}, {Nb: 19}},
		},
		{
			[]Box{{Nb: 3}, {Nb: 15}, {Nb: 0}, {Nb: 2}, {Nb: 22}},
			[]Box{{Nb: 9}, {Nb: 18}, {Nb: 13}, {Nb: 17}, {Nb: 5}},
			[]Box{{Nb: 19}, {Nb: 8}, {Nb: 7}, {Nb: 25}, {Nb: 23}},
			[]Box{{Nb: 20}, {Nb: 11}, {Nb: 10}, {Nb: 24}, {Nb: 4}},
			[]Box{{Nb: 14}, {Nb: 21}, {Nb: 16}, {Nb: 12}, {Nb: 6}},
		},
		{
			[]Box{{Nb: 14}, {Nb: 21}, {Nb: 17}, {Nb: 24}, {Nb: 4}},
			[]Box{{Nb: 10}, {Nb: 16}, {Nb: 15}, {Nb: 9}, {Nb: 19}},
			[]Box{{Nb: 18}, {Nb: 8}, {Nb: 23}, {Nb: 26}, {Nb: 20}},
			[]Box{{Nb: 22}, {Nb: 11}, {Nb: 13}, {Nb: 6}, {Nb: 5}},
			[]Box{{Nb: 2}, {Nb: 0}, {Nb: 12}, {Nb: 3}, {Nb: 7}},
		},
	})
	fmt.Printf("Boards: %v\n", boards)
	fmt.Printf("Numbers: %v\n", numbers)
}

func TestBoardLine_Mark(t *testing.T) {
	tests := []struct {
		desc           string
		bl             BoardRow
		nb             int
		expectedBingo  bool
		expectedMarked []bool
	}{
		{
			desc:           "marked first number, no bingo",
			bl:             BoardRow{{Nb: 22}, {Nb: 13}, {Nb: 17}, {Nb: 11}, {Nb: 0}},
			nb:             22,
			expectedMarked: []bool{true, false, false, false, false},
		},
		{
			desc: "marked first number, bingo",
			bl: BoardRow{
				{Nb: 22},
				{Nb: 13, Marked: true},
				{Nb: 17, Marked: true},
				{Nb: 11, Marked: true},
				{Nb: 0, Marked: true},
			},
			nb:             22,
			expectedBingo:  true,
			expectedMarked: []bool{true, true, true, true, true},
		},
		{
			desc:           "does not mark anything number, no bingo",
			bl:             BoardRow{{Nb: 22}, {Nb: 13}, {Nb: 17}, {Nb: 11}, {Nb: 0}},
			nb:             42,
			expectedBingo:  false,
			expectedMarked: []bool{false, false, false, false, false},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			gotBingo := test.bl.Mark(test.nb)
			assert.Equal(t, gotBingo, test.expectedBingo)
			for i, box := range test.bl {
				assert.Equal(t, box.Marked, test.expectedMarked[i])
			}
		})
	}
}
