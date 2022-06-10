namespace Leetcode.FindMinimumInRotatedSortedArray;

public class Solution {
    public int FindMin(int[] nums) {
        if (nums.Length == 1)
            return nums[0];

        if (nums.Length == 2)
            return nums[0] > nums[1] ? nums[1] : nums[0];

        if (nums[0] < nums[^1])
            return nums[0];

        var left = 0;
        var right = nums.Length - 1;
        while (right >= left) {
            var mid = left + ((right - left) / 2);

            if (nums[mid] > nums[mid + 1]) {
                //        ⌄ mid
                // [5, 6, 7, 2, 3]
                //           ^ mid + 1
                return nums[mid + 1];
            }

            if (nums[mid] < nums[mid - 1]) {
                //        ⌄ mid
                // [6, 7, 2, 3, 5]
                //     ^ mid - 1
                return nums[mid];
            }

            if (nums[mid] > nums[0]) {
                //        ⌄ mid
                // [4, 5, 6, 7, 2, 3]
                // restrain search to the right
                // eg. [7, 2, 3]
                left = mid + 1;
            } else {
                //                ⌄ mid
                // [9, 20, 21, 2, 3, 4, 5, 6, 7]
                // restrain search to the left
                // eg. [9, 20, 21, 2, 3]
                right = mid - 1;
            }
        }

        return -1;
    }
}
