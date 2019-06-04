package problem

func invertTree(root *TreeNode) *TreeNode {
	if root == nil  {
		return root
	}
	root.Right,root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}
