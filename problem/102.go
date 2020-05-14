package problem

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

 

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	level := make([]*TreeNode, 0)
	level = append(level, root)
	nodeCount := len(level)
	for len(level) > 0 {
		levelSlice := make([]int, 0)
		for nodeCount > 0  {
			top := level[0]
			levelSlice = append(levelSlice, top.Val)
			if top.Left != nil {
				level = append(level, top.Left)
			}
			if top.Right != nil {
				level = append(level, top.Right)
			}
			level = level[1:]
			nodeCount --
		}
		ret = append(ret, levelSlice)
		nodeCount = len(level)
	}
	return ret
}