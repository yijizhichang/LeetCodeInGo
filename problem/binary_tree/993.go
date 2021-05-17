package binary_tree

/*
在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。

如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。

我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。

只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。

 

示例 1：


输入：root = [1,2,3,4], x = 4, y = 3
输出：false

 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// bfs
// 旨在找到x节点和y节点
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	update := func(node, parentNode *TreeNode, depth int) {
		if node.Val == x {
			xParent, xDepth, xFound = parentNode, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parentNode, depth, true
		}
	}

	type pair struct {
		node  *TreeNode
		depth int
	}

	q := []pair{{root, 0}}
	update(root, nil, 0)
	for len(q) > 0 && (!xFound || !yFound) {
		node, depth := q[0].node, q[0].depth
		q = q[1:]
		if node.Left != nil {
			q = append(q, pair{node.Left, depth + 1})
			update(node.Left, node, depth+1)
		}
		if node.Right != nil {
			q = append(q, pair{node.Right, depth + 1})
			update(node.Right, node, depth+1)
		}
	}

	return xDepth == yDepth && xParent != yParent
}

// dfs
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool
	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}

		// 如果两个节点都找到了，就可以提前退出遍历
		// 即使不提前退出，对最坏情况下的时间复杂度也不会有影响

		if xFound && yFound {
			return
		}

		// left
		dfs(node.Left, node, depth+1)

		if xFound && yFound {
			return
		}

		// right
		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 0)
	return xDepth == yDepth && xParent != yParent
}