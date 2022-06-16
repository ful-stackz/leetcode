package leetcode

import "sort"

func longestStrChain(words []string) int {
	// sort words in ascending order by len
	sort.Slice(words, func(i, j int) bool {
		return len(words[j]) > len(words[i])
	})

	longest := 0
	dp := make(map[string]int, 0)
	for _, word := range words {
		dp[word] = 1

		for i := range word {
			pred := word[:i] + word[i+1:]
			if dp[pred] != 0 {
				dp[word] = max(dp[word], dp[pred]+1)
			}
		}

		longest = max(longest, dp[word])
	}
	return longest
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
