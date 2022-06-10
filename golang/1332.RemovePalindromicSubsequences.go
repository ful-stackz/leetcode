package leetcode

func removePalindromeSub(s string) int {
	if len(s) == 1 || reverse(s) == s {
		return 1
	}
	return 2
}

func reverse(s string) string {
	rev := make([]rune, len(s))
	for i, v := range s {
		rev[len(s)-i-1] = v
	}
	return string(rev)
}
