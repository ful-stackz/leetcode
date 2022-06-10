package leetcode

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	row := 0
	goingDown := false

	for i := 0; i < len(s); i++ {
		rows[row].WriteByte(s[i])

		if row == 0 || row == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			row++
		} else {
			row--
		}
	}

	var result strings.Builder
	for _, b := range rows {
		result.WriteString(b.String())
	}

	return result.String()
}
