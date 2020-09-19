package problem

/*
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	tmp := make([]*TreeNode, 0)
	now := root
	for len(tmp) > 0 || now != nil {
		for now != nil {
			tmp = append(tmp, now)
			now = now.Left
		}
		now = tmp[len(tmp)-1]
		res = append(res, now.Val)
		tmp = tmp[:len(tmp)-1]
		now = now.Right
	}
	return res
}
