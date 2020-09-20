package leetcode

func inorderTraversal(root *TreeNode) []int {
	var vals []int

	if root == nil {
		return vals
	}

	vals = append(vals, inorderTraversal(root.Left)...)
	vals = append(vals, root.Val)
	vals = append(vals, inorderTraversal(root.Right)...)

	return vals
}
