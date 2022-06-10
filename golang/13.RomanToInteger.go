package leetcode

func romanToInt(s string) int {
	romanNumerals := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	number := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanNumerals[s[i:i+2]] != 0 {
			number += romanNumerals[s[i:i+2]]
			i++
			continue
		}

		number += romanNumerals[s[i:i+1]]
	}

	return number
}
