package sudoku_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/sudoku"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenAndcheck(t *testing.T) {

	Convey("TestGenAndcheck", t, func() {
		sg := &sudoku.SukudoGen{40}
		result, answer := sg.GenSukudo()

		sudoku.Print(result)

		skd := sudoku.NewSukudo()
		skd.ResultIn(result)

		success, _ := skd.Solve()
		So(success, ShouldBeTrue)

		So(answer, ShouldResemble, skd.ResultOut())
	})

}

func ExampleSukudoGen_GenSukudo() {
	sg := &sudoku.SukudoGen{40}
	result, _ := sg.GenSukudo()

	skd := sudoku.NewSukudo()
	skd.ResultIn(result)

	success, _ := skd.Solve()
	fmt.Println(success)

	// Output:
	// true
}
