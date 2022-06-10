namespace Leetcode.LinkedListCycle;

public class Solution {
    public bool HasCycle(ListNode head) {
        var visitedNodes = new Dictionary<ListNode, bool>();
        var node = head;
        while (node is not null) {
            if (visitedNodes.ContainsKey(node))
                return true;

            visitedNodes.Add(node, true);
            node = node.next;
        }

        return false;
    }
}

// Leave out if copying to leetcode
public class ListNode {
    public int val;
    public ListNode? next;
    public ListNode(int x) {
        val = x;
        next = null;
    }
}