package sudoku

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Print sudoku puzzle data
func PrintX(result [][]int, max int, boxW int, boxH int) {

	boxSplitMark := "-----"
	for i := 0; i < boxH-1; i++ {
		boxSplitMark += " + -----"
	}

	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			fmt.Print(result[i][j])
			if j < max-1 {
				fmt.Print(" ")
			}
			if j > 0 && j%boxW == boxW-1 && j < max-1 {
				fmt.Print("| ")
			}
		}
		fmt.Println()

		if i > 0 && i%boxH == boxH-1 && i < max-1 {
			fmt.Println(boxSplitMark)
		}
	}

}

// PrintToString return sudoku 9x9 formt print string
func PrintToString(result [9][9]int) string {
	buf := bytes.NewBuffer(make([]byte, 0))
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fprint(buf, result[i][j])
			if j < 8 {
				fmt.Fprint(buf, " ")
			}
			if j > 0 && j%3 == 2 && j < 8 {
				fmt.Fprint(buf, "| ")
			}
		}
		fmt.Fprintln(buf)

		if i > 0 && i%3 == 2 && i < 8 {
			fmt.Fprintln(buf, "----- + ----- + -----")
		}
	}
	return buf.String()
}

// ParseString to parse string number split by space into [9][9]int format
func ParseString(s string) ([9][9]int, error) {
	var ret [9][9]int

	buf := bytes.NewBufferString(s)

	scanner := bufio.NewScanner(buf)
	if err := scanner.Err(); err != nil {
		return ret, err
	}

	rowid := 0
	for scanner.Scan() {
		strSlice := splitString(scanner.Text())

		for i, singleStr := range strSlice {
			intVal, err := strconv.Atoi(singleStr)
			if err != nil {
				return ret, fmt.Errorf("with a ono integer value '%s'", singleStr)
			}
			if intVal < 0 || intVal > 9 {
				return ret, fmt.Errorf("with invalid integer value '%s', should between 0~9", singleStr)
			}
			if i > 8 {
				break // ignore the rest of the values
			}
			ret[rowid][i] = intVal

		}
		rowid++

		if rowid > 8 {
			break // ignore the rest of the values
		}

	}

	return ret, nil
}

func splitString(str string) []string {
	return strings.Fields(str)
}

// ReadData convert data [][]int to [9][9]int
func ReadData(data [][]int) [9][9]int {
	origin := [9][9]int{}
	if len(data) < 9 {
		return origin
	}
	for i := 0; i < 9; i++ {
		if len(data[i]) < 9 {
			return origin
		}
		for j := 0; j < 9; j++ {
			origin[i][j] = data[i][j]
		}
	}
	return origin
}

func ReadData2(data [9][9]int) [][]int {
	ret := make([][]int, 9)
	for i := 0; i < 9; i++ {
		ret[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			ret[i][j] = data[i][j]
		}
	}
	return ret
}

func Init2dimArray(max int) [][]int {
	ret := make([][]int, max)
	for i := 0; i < max; i++ {
		ret[i] = make([]int, max)
	}
	return ret
}
