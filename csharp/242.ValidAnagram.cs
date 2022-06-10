namespace Leetcode.ValidAnagram;

public class Solution {
    public bool IsAnagram(string s, string t) =>
      s.Length == t.Length &&
      s
        .GroupBy(letter => letter)
        .All(group => s.Count(letter => letter == group.Key) == t.Count(letter => letter == group.Key));
}
