package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	root.Left = root.Right
	root.Right = left

	if root.Left != nil {
		invertTree(root.Left)
	}

	if root.Right != nil {
		invertTree(root.Right)
	}

	return root
}

// Leave out if copying to leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
