package binary_tree

/*
给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
*/

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {
		return rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high) + root.Val
	} else if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	} else if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	} else {
		return 0
	}
}
