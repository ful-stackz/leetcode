package leetcode

func longestCommonPrefix(strs []string) string {
	prefix := []rune(strs[0])
	length := len(strs[0])

	for i := 1; i < len(strs); i++ {
		tLength := 0
		for y, r := range strs[i] {
			if y >= len(prefix) || prefix[y] != r {
				break
			}
			tLength++
		}
		length = min(length, tLength)
		if length == 0 {
			break
		}
	}

	return string(prefix[:length])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
