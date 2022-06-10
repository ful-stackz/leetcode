package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make(map[byte]int, len(s))
	for i := range s {
		letters[s[i]]++
		letters[t[i]]--
	}
	for _, v := range letters {
		if v != 0 {
			return false
		}
	}
	return true
}
