package problem

import (
	"math"
)

/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

 

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

 

注意：两结点之间的路径长度是以它们之间边的数目表示。

 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	//m := 0
	//max := &m
	max := 0
	depth(root, &max)
	return max
}

func depth(now *TreeNode, max *int) int {
	if now == nil {
		return 0
	}
	lDepth := depth(now.Left, max)
	rDepth := depth(now.Right, max)
	*max = int(math.Max(float64(lDepth)+float64(rDepth), float64(*max)))
	return int(math.Max(float64(lDepth), float64(rDepth))) + 1
}
