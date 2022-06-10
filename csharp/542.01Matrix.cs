namespace Leetcode.ZeroOneMatrix;

public class Solution {
    public int[][] UpdateMatrix(int[][] mat) {
        var updatedMatrix = mat.Select(x => x.Select(_ => 0).ToArray()).ToArray();
        var zeroes = FindZeroes(mat);

        for (int row = 0; row < mat.Length; row += 1) {
            for (int col = 0; col < mat[row].Length; col += 1) {
                // skip zeroes
                if (mat[row][col] == 0) continue;

                var closestDistance = int.MaxValue;
                foreach (var point in zeroes) {
                    var distance = Math.Abs(row - point.row) + Math.Abs(col - point.col);
                    if (distance < closestDistance)
                        closestDistance = distance;
                    if (distance == 1)
                        break;
                }

                updatedMatrix[row][col] = closestDistance;
            }
        }

        return updatedMatrix;
    }

    private List<(int row, int col)> FindZeroes(int[][] mat) {
        var points = new List<(int, int)>();
        for (int row = 0; row < mat.Length; row += 1) {
            for (int col = 0; col < mat[row].Length; col += 1) {
                if (mat[row][col] == 0)
                    points.Add((row, col));
            }
        }

        return points;
    }
}
