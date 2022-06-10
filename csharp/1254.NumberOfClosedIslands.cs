namespace Leetcode.NumberOfClosedIslands;

public class Solution {
    public int ClosedIsland(int[][] grid) {
        bool[][] discoveredGrid = grid.Select(row => row.Select(_ => false).ToArray()).ToArray();
        int closedIslands = 0;

        for (int row = 0; row < grid.Length; row += 1) {
            for (int col = 0; col < grid[0].Length; col += 1) {
                if (IsWaterCell(row, col)) continue;
                if (IsDiscoveredCell(row, col)) continue;

                // if this island is touching bounds then later
                // we'll ignore the entire discovered island
                var islandTouchesBounds = false;

                var points = new Stack<(int row, int col)>();
                points.Push((row, col));
                while (points.Any()) {
                    var point = points.Pop();

                    if (point.row < 0 || point.row >= grid.Length) continue;
                    if (point.col < 0 || point.col >= grid[0].Length) continue;

                    if (IsWaterCell(point.row, point.col)) continue;
                    if (IsDiscoveredCell(point.row, point.col)) continue;

                    islandTouchesBounds =
                        islandTouchesBounds || IsBoundingCell(point.row, point.col);

                    discoveredGrid[point.row][point.col] = true;

                    points.Push((point.row - 1, point.col));
                    points.Push((point.row + 1, point.col));
                    points.Push((point.row, point.col - 1));
                    points.Push((point.row, point.col + 1));
                }

                // at this point we've discovered the entire island area
                if (!islandTouchesBounds)
                    closedIslands += 1;
            }
        }

        return closedIslands;

        bool IsWaterCell(int row, int col) =>
            grid[row][col] == 1;

        bool IsDiscoveredCell(int row, int col) =>
            discoveredGrid[row][col];

        bool IsBoundingCell(int row, int col) =>
            row == 0 || row == grid.Length - 1 ||
            col == 0 || col == grid[0].Length - 1;
    }
}
