namespace Leetcode.LongestSubstringWithoutRepeatingCharacters;

public class Solution {
    public int LengthOfLongestSubstring(string s) {
        var set = Enumerable.Range(0, 256).Select(_ => -1).ToArray();
        var start = 0;
        var maxLen = 0;

        for (int end = 0; end < s.Length; end++) {
            var symbol = (int)s[end];
            if (set[symbol] != -1) {
                // this symbol is already part of the current substring
                // set the start of the substring to the symbol
                // after the one that is a duplicate
                start = Math.Max(start, set[symbol] + 1);
            }

            // set symbol index to last position it is encountered at
            set[symbol] = end;

            maxLen = Math.Max(maxLen, end - start + 1);
        }

        return maxLen;
    }
}
