namespace Leetcode.ImplementStrStr;

public class Solution {
    public int StrStr(string haystack, string needle) {
        if (needle == string.Empty) return 0;

        for (int i = 0; i <= haystack.Length - needle.Length; i++) {
			if (haystack[i..(i+needle.Length)] == needle)
				return i;
        }

        return -1;
    }
}
