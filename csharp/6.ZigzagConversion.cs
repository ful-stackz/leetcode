namespace Leetcode.ZigzagConversion;

using System.Text;

public class Solution {
    public string Convert(string s, int numRows) {
        if (numRows == 1)
            return s;

        var rows = Enumerable.Range(0, numRows)
            .Select(_ => new StringBuilder())
            .ToArray();

        int currentRow = 0;
        bool goingDown = false;

        foreach (var c in s) {
            rows[currentRow].Append(c);

            if (currentRow == 0 || currentRow == numRows - 1)
                goingDown = !goingDown;

            currentRow += goingDown ? 1 : -1;
        }

        return String.Join("", rows.Select(row => row.ToString()));
    }
}
