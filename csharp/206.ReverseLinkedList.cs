namespace Leetcode.ReverseLinkedList;

public class Solution {
    public ListNode ReverseList(ListNode head) {
        if (head is null) return null;

        var newHead = new ListNode(head.val, next: null);
        var node = head.next;
        while (node is not null) {
            newHead = new ListNode(node.val, newHead);
            node = node.next;
        }

        return newHead;
    }
}

// Leave out if copying to leetcode
public class ListNode {
    public int val;
    public ListNode? next;
    public ListNode(int val = 0, ListNode? next = null) {
        this.val = val;
        this.next = next;
    }
}
