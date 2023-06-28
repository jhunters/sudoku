package sudoku_test

import (
	"testing"

	"github.com/jhunters/sudoku"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenXAndcheck9x9(t *testing.T) {

	Convey("TestGenXAndcheck 9x9", t, func() {
		sg, err := sudoku.NewSukudoGenX(9, 10)
		So(err, ShouldBeNil)

		result, answer, err := sg.GenSukudo()
		So(err, ShouldBeNil)

		skd, err := sudoku.NewSukudoX(9)
		So(err, ShouldBeNil)
		skd.ResultIn(result)

		success, _ := skd.Solve()
		So(success, ShouldBeTrue)

		So(answer, ShouldResemble, skd.ResultOut())
	})

}
func TestGenXAndcheck6x6(t *testing.T) {
	Convey("TestGenXAndcheck 6x6", t, func() {
		sg, err := sudoku.NewSukudoGenX(6, 10)
		So(err, ShouldBeNil)

		result, answer, err := sg.GenSukudo()
		So(err, ShouldBeNil)

		skd, err := sudoku.NewSukudoX(6)
		So(err, ShouldBeNil)
		skd.ResultIn(result)

		success, _ := skd.Solve()
		So(success, ShouldBeTrue)

		So(answer, ShouldResemble, skd.ResultOut())
	})
}

func TestGenXAndcheck4x4(t *testing.T) {
	Convey("TestGenXAndcheck 4x4", t, func() {
		sg, err := sudoku.NewSukudoGenX(4, 2)
		So(err, ShouldBeNil)

		result, answer, err := sg.GenSukudo()
		So(err, ShouldBeNil)

		skd, err := sudoku.NewSukudoX(4)
		So(err, ShouldBeNil)
		skd.ResultIn(result)

		success, _ := skd.Solve()
		So(success, ShouldBeTrue)

		So(answer, ShouldResemble, skd.ResultOut())

		skd.Print()
	})
}

func TestGenXAndcheck4x2(t *testing.T) {
	Convey("TestGenXAndcheck 4x2", t, func() {
		sg, err := sudoku.NewSukudoGenX(8, 2)
		So(err, ShouldBeNil)

		result, answer, err := sg.GenSukudo()
		So(err, ShouldBeNil)

		skd, err := sudoku.NewSukudoX(8)
		So(err, ShouldBeNil)
		skd.ResultIn(result)

		success, _ := skd.Solve()
		So(success, ShouldBeTrue)

		So(answer, ShouldResemble, skd.ResultOut())

		skd.Print()
	})

}
