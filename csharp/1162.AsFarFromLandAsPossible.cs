namespace Leetcode.AsFarFromLandAsPossible;

public class Solution {
    public int MaxDistance(int[][] grid) {
        var dgrid = new int[grid.Length, grid.Length];
        for (int r = 0; r < grid.Length; r++)
        {
            for (int c = 0; c < grid.Length; c++)
            {
                dgrid[r, c] = (grid[r][c] == 1) ? 0 : int.MaxValue;
            }
        }

        // top, left to bottom, right
        for (int r = 0; r < grid.Length; r++)
        {
            for (int c = 0; c < grid.Length; c++)
            {
                if (grid[r][c] == 0)
                {
                    if (c > 0 && dgrid[r, c - 1] != int.MaxValue)
                    {
                        dgrid[r, c] = Math.Min(dgrid[r, c], 1 + dgrid[r, c - 1]);
                    }

                    if (r > 0 && dgrid[r - 1, c] != int.MaxValue)
                    {
                        dgrid[r, c] = Math.Min(dgrid[r, c], 1 + dgrid[r - 1, c]);
                    }
                }
            }
        }

        int result = 0;

        // bottom, right to top, left
        for (int r = grid.Length - 1; r >= 0; r--)
        {
            for (int c = grid.Length - 1; c >= 0; c--)
            {
                if (grid[r][c] == 0)
                {
                    if (c < grid.Length - 1 && dgrid[r, c + 1] != int.MaxValue)
                    {
                        dgrid[r, c] = Math.Min(dgrid[r, c], 1 + dgrid[r, c + 1]);
                    }

                    if (r < grid.Length - 1 && dgrid[r + 1, c] != int.MaxValue)
                    {
                        dgrid[r, c] = Math.Min(dgrid[r, c], 1 + dgrid[r + 1, c]);
                    }

                    if (dgrid[r, c] != int.MaxValue && dgrid[r, c] > result)
                    {
                        result = dgrid[r, c];
                    }
                }
            }
         }

        return (result == 0) ? -1: result;
    }
}
