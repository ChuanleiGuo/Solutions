package leetcode

func postorderTraversal(root *TreeNode) []int {
	var vals []int

	if root == nil {
		return vals
	}

	vals = append(vals, postorderTraversal(root.Left)...)
	vals = append(vals, postorderTraversal(root.Right)...)
	vals = append(vals, root.Val)

	return vals
}
