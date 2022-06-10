package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	revs1 := mapToInt(strings.Split(version1, "."))
	revs2 := mapToInt(strings.Split(version2, "."))
	for i := 0; i < max(len(revs1), len(revs2)); i++ {
		rev1 := 0
		if i < len(revs1) {
			rev1 = revs1[i]
		}

		rev2 := 0
		if i < len(revs2) {
			rev2 = revs2[i]
		}

		if rev1 > rev2 {
			return 1
		} else if rev2 > rev1 {
			return -1
		}
	}

	return 0
}

func mapToInt(src []string) []int {
	result := make([]int, len(src))
	for i, v := range src {
		result[i] = toInt(v)
	}
	return result
}

func toInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
