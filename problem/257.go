package problem

import (
	"strconv"
	"strings"
)

/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

*/

var res []string

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res = []string{}
	tmp := make([]string, 0)
	helper(root, tmp)
	return res
}

func helper(root *TreeNode, now []string) {
	if root.Left != nil {
		helper(root.Left, append(now, strconv.Itoa(root.Val)))
	}
	if root.Right != nil {
		helper(root.Right, append(now, strconv.Itoa(root.Val)))
	}
	now = append(now, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		res = append(res, strings.Join(now, "->"))
		return
	}
}
