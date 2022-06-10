package leetcode

var LAND_CELL byte = '1'
var VISITED_CELL byte = '2'

func numIslands(grid [][]byte) int {
	islands := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != LAND_CELL {
				continue
			}

			traverse(grid, row, col)
			islands += 1
		}
	}

	return islands
}

func traverse(grid [][]byte, row int, col int) {
	if row < 0 ||
		col < 0 ||
		row >= len(grid) ||
		col >= len(grid[row]) ||
		grid[row][col] != LAND_CELL {
		return
	}

	// mark as traversed
	grid[row][col] = VISITED_CELL

	// traverse neighbors (vert + horiz)
	traverse(grid, row+1, col)
	traverse(grid, row-1, col)
	traverse(grid, row, col+1)
	traverse(grid, row, col-1)

	// return grid
}
