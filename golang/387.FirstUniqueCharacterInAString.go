package leetcode

func firstUniqChar(s string) int {
	chars := make(map[rune]int, len(s))
	for _, r := range s {
		chars[r]++
	}
	for i, r := range s {
		if chars[r] == 1 {
			return i
		}
	}
	return -1
}
