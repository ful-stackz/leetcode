package leetcode

func lengthOfLongestSubstring(s string) int {
	set := make([]int, 256)
	start := 0
	maxLen := 0

	for end := 0; end < len(s); end++ {
		symbol := int(s[end])
		if set[symbol] != 0 {
			start = max(start, set[symbol])
		}

		set[symbol] = end + 1
		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
