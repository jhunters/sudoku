package sudoku_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/sudoku"
	. "github.com/smartystreets/goconvey/convey"
)

// TestSuccess
func TestCheckResultSuccess(t *testing.T) {
	sdk, _ := sudoku.NewSudokuX(9)

	origin := [][]int{{8, 9, 4, 1, 3, 6, 5, 2, 7}, {5, 1, 3, 2, 9, 7, 8, 6, 4},
		{7, 6, 2, 5, 4, 8, 9, 3, 1}, {9, 4, 8, 7, 5, 3, 2, 1, 6}, {6, 2, 5, 4, 8, 1, 3, 7, 9},
		{1, 3, 7, 9, 6, 2, 4, 5, 8}, {3, 5, 9, 6, 1, 4, 7, 8, 2}, {4, 7, 6, 8, 2, 5, 1, 9, 3}, {2, 8, 1, 3, 7, 9, 6, 4, 5},
	}

	sdk.ResultIn(origin)

	if !sdk.Success() {
		t.Error("Check success failed")
	}
}

// TestSolveWithInvalidateData
func TestSolveWithInvalidateData(t *testing.T) {
	sdk, _ := sudoku.NewSudokuX(9)

	origin := [][]int{{8, 9, 0, 1, 0, 0, 0, 2, 0}, {0, 0, 0, 2, 9, 7, 8, 6, 4},
		{0, 0, 0, 5, 4, 8, 9, 3, 1}, {9, 4, 8, 7, 5, 3, 2, 6, 6}, {6, 2, 5, 4, 8, 1, 3, 7, 9},
		{1, 3, 7, 9, 6, 2, 4, 5, 8}, {3, 5, 9, 6, 1, 4, 7, 8, 2}, {4, 7, 6, 8, 2, 5, 1, 9, 3}, {2, 8, 1, 3, 7, 9, 6, 4, 5},
	}

	sdk.ResultIn(origin)

	validate, _ := sdk.Solve()

	if validate {
		t.Error("validate should return false")
	}
}

func TestSolveSuccess(t *testing.T) {
	Convey("Test solve success", t, func() {
		sdk, _ := sudoku.NewSudokuX(9)
		origin := [][]int{{6, 0, 0, 0, 0, 2, 0, 0, 0}, {0, 0, 1, 0, 0, 7, 0, 0, 2},
			{0, 3, 4, 9, 0, 0, 0, 0, 0}, {8, 6, 0, 0, 5, 0, 0, 4, 0}, {1, 0, 0, 0, 0, 0, 0, 0, 6},
			{0, 0, 9, 7, 0, 0, 8, 0, 5}, {0, 0, 0, 0, 2, 0, 9, 6, 0}, {0, 0, 0, 0, 0, 1, 0, 0, 4}, {4, 0, 0, 0, 0, 5, 0, 8, 0},
		}

		expect := [][]int{{6, 7, 8, 5, 3, 2, 4, 9, 1}, {9, 5, 1, 8, 4, 7, 6, 3, 2}, {2, 3, 4, 9, 1, 6, 7, 5, 8},
			{8, 6, 7, 1, 5, 3, 2, 4, 9}, {1, 4, 5, 2, 8, 9, 3, 7, 6}, {3, 2, 9, 7, 6, 4, 8, 1, 5}, {5, 1, 3, 4, 2, 8, 9, 6, 7},
			{7, 8, 6, 3, 9, 1, 5, 2, 4}, {4, 9, 2, 6, 7, 5, 1, 8, 3}}

		sdk.ResultIn(origin)

		result, _ := sdk.Solve()
		So(result, ShouldBeTrue)

		So(sdk.Finished(), ShouldBeTrue)

		data := sdk.ResultOut()
		So(data, ShouldResemble, expect)
	})

}

func TestFromStringSolveSuccess(t *testing.T) {
	Convey("Test solve success", t, func() {
		sdk, _ := sudoku.NewSudokuX(9)
		originStr := "600002000001007002034900000860050040100000006009700805000020960000001004400005080"

		expect := [][]int{{6, 7, 8, 5, 3, 2, 4, 9, 1}, {9, 5, 1, 8, 4, 7, 6, 3, 2}, {2, 3, 4, 9, 1, 6, 7, 5, 8},
			{8, 6, 7, 1, 5, 3, 2, 4, 9}, {1, 4, 5, 2, 8, 9, 3, 7, 6}, {3, 2, 9, 7, 6, 4, 8, 1, 5}, {5, 1, 3, 4, 2, 8, 9, 6, 7},
			{7, 8, 6, 3, 9, 1, 5, 2, 4}, {4, 9, 2, 6, 7, 5, 1, 8, 3}}

		sdk.ResultInFromString(originStr)

		result, _ := sdk.Solve()
		So(result, ShouldBeTrue)

		So(sdk.Finished(), ShouldBeTrue)

		data := sdk.ResultOut()
		So(data, ShouldResemble, expect)

		sdk.Print()
	})

}

func ExampleSudokuX_Solve() {

	data := `6 0 0 0 0 2 0 0 0
	0 0 1 0 0 7 0 0 2
	0 3 4 9 0 0 0 0 0
	8 6 0 0 5 0 0 4 0
	1 0 0 0 0 0 0 0 6
	0 0 9 7 0 0 8 0 5
	0 0 0 0 2 0 9 6 0
	0 0 0 0 0 1 0 0 4
	4 0 0 0 0 5 0 8 0`

	sdk, _ := sudoku.NewSudokuX(9)
	origin, err := sudoku.ParseString(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	sdk.ResultIn(sudoku.ReadData2(origin))

	result, _ := sdk.Solve()
	fmt.Println(result)

	sdk.Print()

	// Output:
	// true
	// 6 7 8 | 5 3 2 | 4 9 1
	// 9 5 1 | 8 4 7 | 6 3 2
	// 2 3 4 | 9 1 6 | 7 5 8
	// ----- + ----- + -----
	// 8 6 7 | 1 5 3 | 2 4 9
	// 1 4 5 | 2 8 9 | 3 7 6
	// 3 2 9 | 7 6 4 | 8 1 5
	// ----- + ----- + -----
	// 5 1 3 | 4 2 8 | 9 6 7
	// 7 8 6 | 3 9 1 | 5 2 4
	// 4 9 2 | 6 7 5 | 1 8 3
}

func ExampleSudokuX_ResultInFromString() {

	data := "600002000001007002034900000860050040100000006009700805000020960000001004400005080"
	sdk, _ := sudoku.NewSudokuX(9)
	sdk.ResultInFromString(data)

	result, _ := sdk.Solve()
	fmt.Println(result)

	sdk.Print()

	// Output:
	// true
	// 6 7 8 | 5 3 2 | 4 9 1
	// 9 5 1 | 8 4 7 | 6 3 2
	// 2 3 4 | 9 1 6 | 7 5 8
	// ----- + ----- + -----
	// 8 6 7 | 1 5 3 | 2 4 9
	// 1 4 5 | 2 8 9 | 3 7 6
	// 3 2 9 | 7 6 4 | 8 1 5
	// ----- + ----- + -----
	// 5 1 3 | 4 2 8 | 9 6 7
	// 7 8 6 | 3 9 1 | 5 2 4
	// 4 9 2 | 6 7 5 | 1 8 3
}
