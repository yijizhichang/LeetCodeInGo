package problem

/*
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	left := 0
	right := len(nums) - 1
	return helper(left, right, nums)
}

func helper(left, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	this := &TreeNode{
		Val:   nums[mid],
		Left:  helper(left, mid-1, nums),
		Right: helper(mid+1, right, nums),
	}
	return this
}
