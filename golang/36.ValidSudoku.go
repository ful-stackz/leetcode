package leetcode

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		// check row
		if isValidRow(board, row) == false {
			return false
		}

		for col := 0; (row == 0 || row%3 == 0) && col < 9; col++ {
			// check column
			if row == 0 && isValidCol(board, col) == false {
				return false
			}

			// check square
			if col%3 == 0 && isValidSquare(board, row, col) == false {
				return false
			}
		}
	}

	return true
}

func isValidRow(board [][]byte, row int) bool {
	hasht := make(map[byte]bool, 9)
	for _, cell := range board[row] {
		if cell == '.' {
			continue
		}

		if hasht[cell] {
			// found a repeating value
			return false
		}
		hasht[cell] = true
	}
	return true
}

func isValidCol(board [][]byte, col int) bool {
	hasht := make(map[byte]bool, 9)
	for _, row := range board {
		if row[col] == '.' {
			continue
		}

		if hasht[row[col]] {
			// found a repeating value
			return false
		}

		hasht[row[col]] = true
	}
	return true
}

func isValidSquare(board [][]byte, row int, col int) bool {
	// row and col wil be [0, 3, 6]
	hasht := make(map[byte]bool, 9)
	for y := row; y < row+3; y++ {
		for x := col; x < col+3; x++ {
			if board[y][x] == '.' {
				continue
			}

			if hasht[board[y][x]] {
				return false
			}
			hasht[board[y][x]] = true
		}
	}
	return true
}
