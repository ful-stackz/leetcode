namespace Leetcode.MergeIntervals;

public class Solution {
    public int[][] Merge(int[][] intervals) {
        var result = new LinkedList<int[]>();
        foreach (var interval in intervals.OrderBy(i => i[0])) {
            if (result.Count == 0 || result.Last()[1] < interval[0]) {
                result.AddLast(interval);
            } else {
                result.Last()[1] = Math.Max(result.Last()[1], interval[1]);
            }
        }

        return result.ToArray();
    }
}
