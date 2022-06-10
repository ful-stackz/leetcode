package leetcode

import "fmt"

func isValidSudoku(board [][]byte) bool {
	positions := make(map[string]bool, 0)
	for row, cols := range board {
		for col, val := range cols {
			if val == '.' {
				continue
			}

			inrow := fmt.Sprintf("%d in row %d", val, row)
			incol := fmt.Sprintf("%d in col %d", val, col)
			insqr := fmt.Sprintf("%d in sqr %dx%d", val, row/3, col/3)

			if positions[inrow] || positions[incol] || positions[insqr] {
				return false
			}

			positions[inrow] = true
			positions[incol] = true
			positions[insqr] = true
		}
	}

	return true
}
