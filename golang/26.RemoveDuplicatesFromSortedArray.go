package leetcode

func removeDuplicates(nums []int) int {
	encounters := make(map[int]bool, len(nums))
	placeIndex := 0
	for _, val := range nums {
		if encounters[val] {
			continue
		}

		encounters[val] = true
		nums[placeIndex] = val
		placeIndex++
	}

	return placeIndex
}
