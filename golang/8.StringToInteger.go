package leetcode

import (
	"math"
)

func myAtoi(s string) int {
	isNegative := false
	digits := make([]int, 0, len(s))
	gotDigits := false
	gotNonZeroDigits := false
	for _, v := range s {
		if !gotDigits {
			if v == ' ' {
				// allow leading whitespace
				continue
			}

			if v == '-' || v == '+' {
				isNegative = v == '-'
				// the sign must denote the start of the number, if the next
				// character is not a digit the parsing should stop
				gotDigits = true
				continue
			}

			if v < '0' || v > '9' {
				// non-digit character encountered before
				// reading atleast 1 digit,
				// stop parsing
				break
			}
		} else if v < '0' || v > '9' {
			// non-digit character encountered after
			// reading atleast 1 digit,
			// stop parsing
			break
		}

		// v is a digit character
		gotDigits = true
		if v == '0' && !gotNonZeroDigits {
			// ignore leading zeros
			continue
		}

		digits = append(digits, toInt(v))
		gotNonZeroDigits = true
	}

	if !gotDigits {
		return 0
	}

	minNumber := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 8}
	minNumberInt := -2147483648
	maxNumber := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
	maxNumberInt := 2147483647

	if len(digits) > len(minNumber) {
		// if we have more digits than the min/max number then
		// it is certain that the number is out of bounds
		if isNegative {
			return minNumberInt
		} else {
			return maxNumberInt
		}
	}

	// if the parsed number has the same amount of digits as the min/max
	// then flag that it is possible that it could overflow
	possibleMinOverflow := isNegative && len(digits) == len(minNumber)
	possibleMaxOverflow := !isNegative && len(digits) == len(maxNumber)

	// keep the value of the actual number here
	number := 0

	for i, d := range digits {
		if possibleMinOverflow && isNegative {
			if digits[i] < minNumber[i] {
				// the current digit is less than the corresponding digit
				// of the min number; since we are going over the digits
				// left-to-right, any further digit that is > than the
				// corresponding min digit is not of significance, eg.
				// min = -2147483648
				// eg. = -2139999999
				//          ^ this is less, so the rest is not significant
				possibleMinOverflow = false
			}
			if digits[i] > minNumber[i] {
				return minNumberInt
			}
		}

		if possibleMaxOverflow && !isNegative {
			if digits[i] < maxNumber[i] {
				// the current digit is less than the corresponding digit
				// of the max number; since we are going over the digits
				// left-to-right, any further digit that is > than the
				// corresponding max digit is not of significance, eg.
				// max = 2147483647
				// eg. = 2139999999
				//         ^ this is less, so the rest is not significant
				possibleMaxOverflow = false
			}
			if digits[i] > maxNumber[i] {
				return maxNumberInt
			}
		}

		number += int(math.Pow10(len(digits)-i-1)) * d
	}

	if isNegative {
		return -number
	}

	return number
}

func toInt(r rune) int {
	return int(r - '0')
}
