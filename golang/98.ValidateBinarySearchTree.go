package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBSTInternal(root, math.MinInt, math.MaxInt)
}

func isValidBSTInternal(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isValidBSTInternal(node.Left, min, node.Val) &&
		isValidBSTInternal(node.Right, node.Val, max)
}

// Leave out if coptying to leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
