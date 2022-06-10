package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	steps := 0
	right := head
	target := head
	left := head
	for right != nil {
		if steps >= n {
			target = target.Next
		}

		if steps-1 >= n {
			left = left.Next
		}

		steps++

		if right.Next == nil {
			// reached the end
			if steps == n {
				// need to remove the head / first node
				head = target.Next
			} else {
				// need to remove a non-head node
				left.Next = target.Next
			}
		}

		right = right.Next
	}

	return head
}

// Leave out if copying to leetcode
type ListNode struct {
	Val  int
	Next *ListNode
}
