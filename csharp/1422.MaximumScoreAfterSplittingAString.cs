namespace Leetcode.MaximumScoreAfterSplittingAString;

public class Solution {
    public int MaxScore(string s) {
        int maxScore = 0;

        for (int i = 1; i < s.Length; i += 1) {
            var left = s[0..i];
            var right = s[i..];

            var scoreLeft = left.Count(x => x == '0');
            var scoreRight = right.Count(x => x == '1');
            var totalScore = scoreLeft + scoreRight;

            maxScore = totalScore > maxScore ? totalScore : maxScore;
        }

        return maxScore;
    }
}
