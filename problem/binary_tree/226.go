package binary_tree

/*
226. 翻转二叉树
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/
func invertTree(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return
}
