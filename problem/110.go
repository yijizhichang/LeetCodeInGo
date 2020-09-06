package problem

import "math"

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 后序遍历复杂度低
func isBalanced(root *TreeNode) bool {
	return depth(root) >= 0
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left)

	right := depth(root.Right)

	if left == -1 || right == -1 || math.Abs(float64(left)-float64(right)) > float64(1) {
		return -1
	}
	return max(left, right) + 1
}
