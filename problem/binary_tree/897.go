package binary_tree

/*
给你一棵二叉搜索树，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，
使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。


示例 1：
输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	ret := &TreeNode{}
	head := &TreeNode{}
	head = ret

	var helper = func(root *TreeNode) {}
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		ret.Right = root
		ret = ret.Right
		ret.Left = nil
		helper(root.Right)
	}
	helper(root)
	return head.Right
}
