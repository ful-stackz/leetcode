namespace Leetcode.ValidPalindrome;

public class Solution {
    public bool IsPalindrome(string s) {
        var input = string.Join(string.Empty, s.Where(char.IsLetterOrDigit)).ToLower();
        return string.Join(string.Empty, input.Reverse()).Equals(input);
    }
}
