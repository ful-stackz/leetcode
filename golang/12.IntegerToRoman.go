package leetcode

import (
	"sort"
	"strings"
)

func intToRoman(num int) string {
	roman := ""
	tnum := num

	romanNumerals := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	keys := make([]int, 0, len(romanNumerals))
	for k := range romanNumerals {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}

	for _, value := range keys {
		letter := romanNumerals[value]
		count := tnum / value
		if count >= 1 {
			tnum %= value
			roman += strings.Repeat(letter, count)
		}
	}

	return roman
}
