namespace Leetcode.TwoSum;

public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        for (int i = 0; i < nums.Length - 1; i++) {
          for (int y = i + 1; y < nums.Length; y++) {
            if (nums[i] + nums[y] == target)
                return new[] { i, y };
          }
        }

        throw new Exception("Each input has exactly one solution.");
    }
}
