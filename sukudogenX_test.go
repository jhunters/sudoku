package sudoku_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/sudoku"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenXAndcheck(t *testing.T) {

	Convey("TestGenXAndcheck", t, func() {

		sg, err := sudoku.NewSukudoGenX(9, 0)
		So(err, ShouldBeNil)

		result, answer, err := sg.GenSukudo()
		So(err, ShouldBeNil)
		fmt.Println(result, answer)

		rdate := sudoku.ReadData(result)

		// result2 := sudoku.ReadData(result)
		skd2 := sudoku.NewSukudo()
		skd2.ResultIn(rdate)
		skd2.Solve()
		fmt.Println(skd2.ResultOut())

		// skd, err := sudoku.NewSukudoX(9)
		// So(err, ShouldBeNil)
		// skd.ResultIn(result)

		// success, _ := skd.Solve()
		// So(success, ShouldBeTrue)

		// So(answer, ShouldResemble, skd.ResultOut())
	})

}
