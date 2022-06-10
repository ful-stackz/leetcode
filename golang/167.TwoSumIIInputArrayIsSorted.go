package leetcode

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		n1 := numbers[i]
		n2 := target - n1

		// if the required n2 is outside the given numbers
		// continue with the next one
		if n2 < numbers[0] || n2 > numbers[len(numbers)-1] {
			continue
		}

		n2Index := indexOf(&numbers, n2, i)
		if n2Index != -1 {
			return []int{i + 1, n2Index + 1}
		}
	}

	// this point should not be reached as the description
	// states there is exactly one solution
	panic("This point should not be reached")
}

func indexOf(src *[]int, v int, ignoreIndex int) int {
	left := 0
	right := len(*src) - 1
	for right >= left {
		mid := left + ((right - left) / 2)

		if (*src)[mid] == v {
			if mid != ignoreIndex {
				return mid
			}

			if mid-1 >= 0 && (*src)[mid-1] == v {
				return mid - 1
			}

			if mid+1 < len(*src) && (*src)[mid+1] == v {
				return mid + 1
			}

			return -1
		}

		if (*src)[mid] > v {
			right = mid - 1
		}

		if (*src)[mid] < v {
			left = mid + 1
		}
	}

	return -1
}
