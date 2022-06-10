package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := l1
	n2 := l2

	result := &ListNode{}
	resultHead := result

	carry := 0

	for n1 != nil || n2 != nil {
		sum := carry

		if n1 != nil {
			sum += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			sum += n2.Val
			n2 = n2.Next
		}

		carry = sum / 10
		sum %= 10

		result.Val = sum

		if n1 != nil || n2 != nil {
			result.Next = &ListNode{0, nil}
			result = result.Next
		}
	}

	if carry != 0 {
		result.Next = &ListNode{carry, nil}
	}

	return resultHead
}

// Leave out this declaration when copying to leetcode
type ListNode struct {
	Val  int
	Next *ListNode
}
