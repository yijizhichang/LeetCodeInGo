package problem

/*

计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, false)
}

func dfs(root *TreeNode, isLeft bool) int {
	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}
	left, right := 0, 0
	if root.Left != nil {
		left = dfs(root.Left, true)
	}
	if root.Right != nil {
		right = dfs(root.Right, false)
	}
	return left + right
}
