namespace Leetcode.StreamOfCharacters;

using System.Text;

public class TrieNode {
    public TrieNode[] Children { get; } = new TrieNode[26];
    public bool IsWord { get; set; }
}

public class StreamChecker {
    private readonly TrieNode _root = new TrieNode();
    private StringBuilder _letters = new StringBuilder();

    public StreamChecker(string[] words) {
        foreach (var word in words)
            AddWatchedWord(word);
    }

    public bool Query(char letter) {
        _letters.Append(letter);

        var node = _root;
        for (int i = _letters.Length - 1; i >= 0; i--) {
            var index = _letters[i] - 'a';

            // the current letter is not part of any word, return early
            if (node.Children[index] is null)
                return false;

            // the current letter is the last letter of a watched word
            if (node.Children[index].IsWord)
                return true;

            // switch node to continue checking down the tree
            node = node.Children[index];
        }

        return false;
    }

    private void AddWatchedWord(string word) {
        var node = _root;

        foreach (var letter in word.Reverse()) {
            int index = letter - 'a';
            if (node.Children[index] is null)
                node.Children[index] = new TrieNode();
            node = node.Children[index];
        }

        node.IsWord = true;
    }
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * StreamChecker obj = new StreamChecker(words);
 * bool param_1 = obj.Query(letter);
 */
