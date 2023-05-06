package sudoku_test

import (
	"testing"

	"github.com/jhunters/sudoku"
)

func TestPrint(t *testing.T) {

	expect := [9][9]int{{6, 7, 8, 5, 3, 2, 4, 9, 1}, {9, 5, 1, 8, 4, 7, 6, 3, 2}, {2, 3, 4, 9, 1, 6, 7, 5, 8},
		{8, 6, 7, 1, 5, 3, 2, 4, 9}, {1, 4, 5, 2, 8, 9, 3, 7, 6}, {3, 2, 9, 7, 6, 4, 8, 1, 5}, {5, 1, 3, 4, 2, 8, 9, 6, 7},
		{7, 8, 6, 3, 9, 1, 5, 2, 4}, {4, 9, 2, 6, 7, 5, 1, 8, 3}}

	sudoku.Print(expect)
}
