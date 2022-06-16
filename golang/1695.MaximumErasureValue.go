package leetcode

func maximumUniqueSubarray(nums []int) int {
	used := make(map[int]bool, 0)
	sum, maxSum := 0, 0
	l, r := 0, 0
	for r < len(nums) {
		if used[nums[r]] == false {
			// number has not been used yet
			sum += nums[r]
			maxSum = max(maxSum, sum)
			used[nums[r]] = true
			r++
		} else {
			// subtract left-most number from curr sum
			sum -= nums[l]
			// mark left-most number as no longer used
			used[nums[l]] = false
			// slide window to the right
			l++
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
