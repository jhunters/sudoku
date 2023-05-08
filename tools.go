package sudoku

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Print sudoku puzzle data
func Print(result [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(result[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
			if j > 0 && j%3 == 2 && j < 8 {
				fmt.Print("| ")
			}
		}
		fmt.Println()

		if i > 0 && i%3 == 2 && i < 8 {
			fmt.Println("----- + ----- + -----")
		}
	}

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
