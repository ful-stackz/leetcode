function twoSum(nums: number[], target: number): number[] {
  for (let i = 0; i < nums.length - 1; i += 1) {
    for (let y = i + 1; y < nums.length; y += 1) {
      if (nums[i] + nums[y] === target) return [i, y]
    }
  }
}
