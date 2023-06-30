package sudoku_test

import (
	"testing"

	"github.com/jhunters/sudoku"
)

func TestSolveX(t *testing.T) {

	data := [][]int{{0, 0, 1, 2, 0, 0}, {0, 0, 0, 0, 5, 0}, {5, 0, 0, 1, 0, 6}, {0, 3, 0, 0, 0, 0}, {6, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 4}}
	sdu, _ := sudoku.NewSudokuX(6)
	sdu.ResultIn(data)
	ok, _ := sdu.Solve()
	if !ok {
		t.Error("solve failed")
	}

	sdu.Print()
}
