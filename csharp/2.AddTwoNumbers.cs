namespace Leetcode.AddTwoNumbers;

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
    public ListNode AddTwoNumbers(ListNode l1, ListNode l2) {
        var n1 = l1;
        var n2 = l2;

        var head = new ListNode();
        var node = head;

        int carry = 0;

        while (n1 is not null || n2 is not null) {
            var sum = carry + (n1?.val ?? 0) + (n2?.val ?? 0);
            carry = sum / 10;

            node.val = sum % 10;

            n1 = n1?.next;
            n2 = n2?.next;

            if (n1 is not null || n2 is not null) {
                node.next = new ListNode();
                node = node.next;
            }
        }

        if (carry != 0) {
            node.next = new ListNode(carry);
        }

        return head;
    }
}

// Leave out if copying over to leetcode
public class ListNode {
    public int val;
    public ListNode? next;
    public ListNode(int val = 0, ListNode? next = null) {
        this.val = val;
        this.next = next;
    }
}
