package sudoku_test

import (
	"testing"

	"github.com/jhunters/sudoku"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenXAndcheck9x9(t *testing.T) {

	Convey("TestGenXAndcheck 9x9", t, func() {
		sg, err := sudoku.NewSudokuGenX(9, 10)
		So(err, ShouldBeNil)

		result, answer, err := sg.GenSudoku()
		So(err, ShouldBeNil)

		sdk, err := sudoku.NewSudokuX(9)
		So(err, ShouldBeNil)
		sdk.ResultIn(result)
		sdk.Print()

		success, _ := sdk.Solve()
		So(success, ShouldBeTrue)

		So(answer, ShouldResemble, sdk.ResultOut())

		sdk.Print()
	})

}
func TestGenXAndcheck6x6(t *testing.T) {
	Convey("TestGenXAndcheck 6x6", t, func() {
		sg, err := sudoku.NewSudokuGenX(6, 10)
		So(err, ShouldBeNil)

		result, answer, err := sg.GenSudoku()
		So(err, ShouldBeNil)

		sdk, err := sudoku.NewSudokuX(6)
		So(err, ShouldBeNil)
		sdk.ResultIn(result)

		success, _ := sdk.Solve()
		So(success, ShouldBeTrue)

		So(answer, ShouldResemble, sdk.ResultOut())
	})
}

func TestGenXAndcheck4x4(t *testing.T) {
	Convey("TestGenXAndcheck 4x4", t, func() {
		sg, err := sudoku.NewSudokuGenX(4, 2)
		So(err, ShouldBeNil)

		result, answer, err := sg.GenSudoku()
		So(err, ShouldBeNil)

		sdk, err := sudoku.NewSudokuX(4)
		So(err, ShouldBeNil)
		sdk.ResultIn(result)

		success, _ := sdk.Solve()
		So(success, ShouldBeTrue)

		So(answer, ShouldResemble, sdk.ResultOut())

		sdk.Print()
	})
}

func TestGenXAndcheck4x2(t *testing.T) {
	Convey("TestGenXAndcheck 4x2", t, func() {
		sg, err := sudoku.NewSudokuGenX(8, 2)
		So(err, ShouldBeNil)

		result, answer, err := sg.GenSudoku()
		So(err, ShouldBeNil)

		sdk, err := sudoku.NewSudokuX(8)
		So(err, ShouldBeNil)
		sdk.ResultIn(result)

		success, _ := sdk.Solve()
		So(success, ShouldBeTrue)

		So(answer, ShouldResemble, sdk.ResultOut())

		sdk.Print()
	})

}
