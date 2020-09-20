package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var vals []int

	if root == nil {
		return vals
	}

	vals = append(vals, root.Val)
	vals = append(vals, preorderTraversal(root.Left)...)
	vals = append(vals, preorderTraversal(root.Right)...)

	return vals
}
