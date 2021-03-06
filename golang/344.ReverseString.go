package leetcode

func reverseString(s []byte) {
	if len(s) == 1 {
		return
	}

	for i := 0; i < len(s)/2; i++ {
		l1, l2 := s[i], s[len(s)-1-i]
		s[i], s[len(s)-1-i] = l2, l1
	}
}
