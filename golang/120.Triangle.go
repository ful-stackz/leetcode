package leetcode

func minimumTotal(triangle [][]int) int {
	for row := range triangle {
		if row == 0 {
			// ignore first row
			continue
		}

		for col := range triangle[row] {
			if col == 0 {
				triangle[row][col] += triangle[row-1][col]
			} else if col == len(triangle[row])-1 {
				triangle[row][col] += triangle[row-1][col-1]
			} else {
				triangle[row][col] += min(triangle[row-1][col-1], triangle[row-1][col])
			}
		}
	}

	return arrMin(triangle[len(triangle)-1])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func arrMin(a []int) int {
	r := a[0]
	for i := 1; i < len(a); i++ {
		r = min(r, a[i])
	}
	return r
}
