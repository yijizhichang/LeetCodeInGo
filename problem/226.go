package problem

/*
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
	if root == nil {
		return root
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}
