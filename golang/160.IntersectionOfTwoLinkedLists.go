package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodea := headA
	nodeavals := make(map[*ListNode]bool)
	for nodea != nil {
		nodeavals[nodea] = true
		nodea = nodea.Next
	}

	nodeb := headB
	for nodeb != nil {
		if nodeavals[nodeb] {
			return nodeb
		}
		nodeb = nodeb.Next
	}

	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
