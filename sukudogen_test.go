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

		// sdu, _ := sudoku.NewSukudoX(9)
		// rdate := sudoku.Init2dimArray(9)
		// for i := 0; i < 9; i++ {
		// 	copy(rdate[i], result[i][:])
		// }
		// sdu.ResultIn(rdate)
		// ok, count := sdu.Solve()
		// fmt.Println(ok, count)
		// fmt.Println(sdu.ResultOut())

		skd := sudoku.NewSukudo()
		skd.ResultIn(result)

		success, count := skd.Solve()
		fmt.Println(success, count)
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
