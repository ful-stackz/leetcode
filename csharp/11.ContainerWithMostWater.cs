namespace Leetcode.ContainerWithMostWater;

public class Solution {
    public int MaxArea(int[] height) {
        int left = 0;
        int right = height.Length - 1;
        int maxArea = -1;

        while (right > left) {
            int maxL = height[left];
            int maxR = height[right];
            int area = CalculateArea(left, right);
            if (area > maxArea)
                maxArea = area;

            if (maxL > maxR) {
                // start looking for a new maxR
                while (right > left) {
                    right--;
                    if (height[right] > maxR) break;
                }
            } else {
                // start looking for a new maxL
                while (right > left) {
                    left++;
                    if (height[left] > maxL) break;
                }
            }
        }

        return maxArea;

        int CalculateArea(int p1, int p2) {
            return Math.Min(height[p1], height[p2]) * Math.Abs(p1 - p2);
        }
    }
}
