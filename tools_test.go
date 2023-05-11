package sudoku_test

import (
	"testing"

	"github.com/jhunters/sudoku"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPrint(t *testing.T) {

	expect := [9][9]int{{6, 7, 8, 5, 3, 2, 4, 9, 1}, {9, 5, 1, 8, 4, 7, 6, 3, 2}, {2, 3, 4, 9, 1, 6, 7, 5, 8},
		{8, 6, 7, 1, 5, 3, 2, 4, 9}, {1, 4, 5, 2, 8, 9, 3, 7, 6}, {3, 2, 9, 7, 6, 4, 8, 1, 5}, {5, 1, 3, 4, 2, 8, 9, 6, 7},
		{7, 8, 6, 3, 9, 1, 5, 2, 4}, {4, 9, 2, 6, 7, 5, 1, 8, 3}}

	sudoku.Print(expect)
}
func TestPrintToString(t *testing.T) {

	expect := [9][9]int{{6, 7, 8, 5, 3, 2, 4, 9, 1}, {9, 5, 1, 8, 4, 7, 6, 3, 2}, {2, 3, 4, 9, 1, 6, 7, 5, 8},
		{8, 6, 7, 1, 5, 3, 2, 4, 9}, {1, 4, 5, 2, 8, 9, 3, 7, 6}, {3, 2, 9, 7, 6, 4, 8, 1, 5}, {5, 1, 3, 4, 2, 8, 9, 6, 7},
		{7, 8, 6, 3, 9, 1, 5, 2, 4}, {4, 9, 2, 6, 7, 5, 1, 8, 3}}

	ret := sudoku.PrintToString(expect)
	s := "6 7 8 | 5 3 2 | 4 9 1\n9 5 1 | 8 4 7 | 6 3 2\n2 3 4 | 9 1 6 | 7 5 8\n----- + ----- + -----\n8 6 7 | 1 5 3 | 2 4 9\n1 4 5 | 2 8 9 | 3 7 6\n3 2 9 | 7 6 4 | 8 1 5\n----- + ----- + -----\n5 1 3 | 4 2 8 | 9 6 7\n7 8 6 | 3 9 1 | 5 2 4\n4 9 2 | 6 7 5 | 1 8 3\n"

	Convey("TestPrintToString", t, func() {
		So(ret, ShouldEqual, s)
	})
}

func TestParseString(t *testing.T) {

	Convey("TestParseString ok", t, func() {
		s := `6 7 8 5 3 2 4 9 1
		9 5 1 8 4 7 6 3 2
		2 3 4 9 1 6 7 5 8
		8 6 7 1 5 3 2 4 9 
		1 4 5 2 8 9 3 7 6
		3 2 9 7 6 4 8 1 5
		5 1 3 4 2 8 9 6 7
		7 8 6 3 9 1 5 2 4
		4 9 2 6 7 5 1 8 3`

		data, err := sudoku.ParseString(s)
		So(err, ShouldBeNil)
		So(data[8][8], ShouldEqual, 3)
	})

	Convey("TestParseString not enough data", t, func() {
		s := `6 7 8 5 3 2 
		9 5 1 8 4 7 6 3 2
		2 3 4 9 1 6 7 5 8
		8 6 7 1 5 3 2 4 9 
		1 4 5 2 8 9 
		3 2 9 7 6 4 8 1 5
		5 1 3 4 2 8 9 6 7
		7 8 6 3 9 1 
		4 9 2 6 7 5 1 8 3`

		data, err := sudoku.ParseString(s)
		So(err, ShouldBeNil)
		So(data[0][6], ShouldEqual, 0)
		So(data[8][8], ShouldEqual, 3)
	})

	Convey("TestParseString with more data", t, func() {
		s := `6 7 8 5 3 2 4 9 1 3
		9 5 1 8 4 7 6 3 2 3
		2 3 4 9 1 6 7 5 8 2
		8 6 7 1 5 3 2 4 9 3
		1 4 5 2 8 9 3 7 6 3
		3 2 9 7 6 4 8 1 5 3
		5 1 3 4 2 8 9 6 7 4
		7 8 6 3 9 1 5 2 4 5
		4 9 2 6 7 5 1 8 3 6
		4 9 2 6 7 5 1 8 3 6`

		data, err := sudoku.ParseString(s)
		So(err, ShouldBeNil)
		So(data[8][8], ShouldEqual, 3)
	})

	Convey("TestParseString with invalid integer", t, func() {
		s := `6 7 8 5 34 2 4 9 1 3
		9 5 1 8 4 7 6 3 2 3
		2 3 4 9 1 6 7 5 8 2
		8 6 7 1 5 3 2 4 9 3
		1 4 5 2 8 9 3 7 6 3
		3 2 9 7 6 4 8 1 5 3
		5 1 3 4 2 8 9 6 7 4
		7 8 6 3 9 1 5 2 4 5
		4 9 2 6 7 5 1 8 3 6
		4 9 2 6 7 5 1 8 3 6`

		_, err := sudoku.ParseString(s)
		So(err, ShouldNotBeNil)
	})

	Convey("TestParseString with invalid integer", t, func() {
		s := `6 7 8 5 a 2 4 9 1 3
		9 5 1 8 4 7 6 3 2 3
		2 3 4 9 1 6 7 5 8 2
		8 6 7 1 5 3 2 4 9 3
		1 4 5 2 8 9 3 7 6 3
		3 2 9 7 6 4 8 1 5 3
		5 1 3 4 2 8 9 6 7 4
		7 8 6 3 9 1 5 2 4 5
		4 9 2 6 7 5 1 8 3 6
		4 9 2 6 7 5 1 8 3 6`

		_, err := sudoku.ParseString(s)
		So(err, ShouldNotBeNil)
	})
}
