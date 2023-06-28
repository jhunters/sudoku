package sudoku_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/sudoku"
)

func TestSolve(t *testing.T) {

	data := [6][6]int{{0, 0, 1, 2, 0, 0}, {0, 0, 0, 0, 5, 0}, {5, 0, 0, 1, 0, 6}, {0, 3, 0, 0, 0, 0}, {6, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 4}}
	sdu := sudoku.NewSukudo6x6()
	sdu.ResultIn(data)
	ok, count := sdu.Solve()
	fmt.Println(ok, count)
	sudoku.Print6x6(sdu.ResultOut())
}

func TestSolveX(t *testing.T) {

	data := [][]int{{0, 0, 1, 2, 0, 0}, {0, 0, 0, 0, 5, 0}, {5, 0, 0, 1, 0, 6}, {0, 3, 0, 0, 0, 0}, {6, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 4}}
	sdu, _ := sudoku.NewSukudoX(6)
	sdu.ResultIn(data)
	ok, count := sdu.Solve()
	fmt.Println(ok, count)
	fmt.Println(sdu.ResultOut())
}
