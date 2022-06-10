namespace Leetcode.NumberOfIslands;

public class Solution {
    public int NumIslands(char[][] grid) {
        int islandsCount = 0;
        var discoveryGrid = grid.Select(row => row.Select(_ => false).ToArray()).ToArray();

        for (int row = 0; row < grid.Length; row += 1) {
            for (int col = 0; col < grid[0].Length; col += 1) {
                // is water, continue
                if (grid[row][col] == '0') continue;

                // is land, is discovered, skip
                if (discoveryGrid[row][col]) continue;

                var points = new Stack<(int row, int col)>();
                points.Push((row, col));
                while (points.Any()) {
                    var point = points.Pop();

                    // is water, ignore
                    if (grid[point.row][point.col] == '0') continue;

                    // is already discovered, ignore
                    if (discoveryGrid[point.row][point.col]) continue;

                    // mark as discovered
                    discoveryGrid[point.row][point.col] = true;

                    // add neighbors for traversal
                    if (point.row - 1 >= 0) points.Push((point.row - 1, point.col));
                    if (point.row + 1 < grid.Length) points.Push((point.row + 1, point.col));
                    if (point.col - 1 >= 0) points.Push((point.row, point.col - 1));
                    if (point.col + 1 < grid[0].Length) points.Push((point.row, point.col + 1));
                }

                // at this point the island is traversed
                islandsCount += 1;
            }
        }

        return islandsCount;
    }
}
