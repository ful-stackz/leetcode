package leetcode

func minOperations(nums []int, x int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	maxLength := -1
	currSum := 0
	for l, r := 0, 0; r < len(nums); r++ {
		currSum += nums[r]
		for l <= r && currSum > sum-x {
			currSum -= nums[l]
			l++
		}
		if currSum == sum-x {
			maxLength = max(maxLength, r-l+1)
		}
	}
	if maxLength == -1 {
		return -1
	}
	return len(nums) - maxLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
