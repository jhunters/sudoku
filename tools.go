package sudoku

import (
	"fmt"
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
