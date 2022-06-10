package leetcode

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for x := i + 1; x < len(nums); x++ {
			if nums[i]+nums[x] == target {
				return []int{i, x}
			}
		}
	}

	panic("Each input has exactly one solution. This point should not be reached.")
}
