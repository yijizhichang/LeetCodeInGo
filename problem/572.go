package problem

import (
	"fmt"
	"strings"
)

/*
给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

     3
    / \
   4   5
  / \
 1   2
给定的树 t：

   4
  / \
 1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 2:
给定的树 s：

     3
    / \
   4   5
  / \
 1   2
    /
   0
给定的树 t：

   4
  / \
 1   2

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
1 将前序遍历组成字符串 判断是否包含关系
2 异常情况1：叶子节点缺少 如果主树： 1是根节点 2是左叶子节点 而子树是1是根节点 2是右叶子节点 他们两都是遍历结果都是12 但是他们两是不包含的 所以需要处理空节点问题，因此我这边是对空节点处理成一个字符 注意左右节点使用不同的符号标记
2 异常情况2 如果主树 12是根节点 而子树 2是根节点 他们两的字符串分别是12 和2 但是他们两也不是包含关系的 需要在特殊处理一下节点值的问题，我这边使用的一个括号将值分割开

*/
func isSubtree(s *TreeNode, t *TreeNode) bool {
	var s1 strings.Builder
	preOrder(s, &s1, "")
	var s2 strings.Builder
	preOrder(t, &s2, "")
	return strings.Contains(s1.String(), s2.String())
}
func preOrder(root *TreeNode, res *strings.Builder, tag string) {
	if root == nil {
		res.WriteString(tag)
		return
	}
	res.WriteString(fmt.Sprintf("(%d)", root.Val))
	preOrder(root.Left, res, "L")
	preOrder(root.Right, res, "R")
}

// dfs暴力解开
// 这是一种最朴素的方法 —— DFS 枚举 ss 中的每一个节点，
// 判断这个点的子树是否和 tt 相等。如何判断一个节点的子树是否和 tt 相等呢，我们又需要做一次 DFS 来检查，
// 即让两个指针一开始先指向该节点和 tt 的根，然后「同步移动」两根指针来「同步遍历」这两棵树，判断对应位置是否相等。
/*
func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil {
        return false
    }
    return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func check(a, b *TreeNode) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil {
        return false
    }
    if a.Val == b.Val {
        return check(a.Left, b.Left) && check(a.Right, b.Right)
    }
    return false
}

*/
