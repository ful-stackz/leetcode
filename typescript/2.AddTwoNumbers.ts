/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */
function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
    let n1 = l1
    let n2 = l2

    let head = new ListNode()
    let node = head

    let carry = 0

    while (n1 != null || n2 != null) {
        let sum = carry + (n1?.val ?? 0) + (n2?.val ?? 0)
        carry = Math.floor(sum / 10)
        sum %= 10

        node.val = sum

        n1 = n1?.next
        n2 = n2?.next

        if (n1 != null || n2 != null) {
            node.next = new ListNode()
            node = node.next
        }
    }

    if (carry != 0) {
        node.next = new ListNode(carry, null)
    }

    return head
}

// Leave out if copying over to leetcode
class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}
