package binary_tree

/*
103. 二叉树的锯齿形层序遍历
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := make([][]int, 0)
	nodeList := make([]*TreeNode, 0)
	left := true
	nodeList = append(nodeList, root)
	tmpNodes := make([]*TreeNode, 0)
	thisFloorRet := make([]int, 0)
	for len(nodeList) > 0 {
		if left {
			for _, thisFloorNode := range nodeList {
				thisFloorRet = append(thisFloorRet, thisFloorNode.Val)
			}
		} else {
			for i := len(nodeList) - 1; i >= 0; i-- {
				thisFloorRet = append(thisFloorRet, nodeList[i].Val)
			}
		}

		for _, thisFloor := range nodeList {
			if thisFloor.Left != nil {
				tmpNodes = append(tmpNodes, thisFloor.Left)
			}
			if thisFloor.Right != nil {
				tmpNodes = append(tmpNodes, thisFloor.Right)
			}
		}
		nodeList = []*TreeNode{}
		left = !left
		ret = append(ret, thisFloorRet)
		thisFloorRet = []int{}
		nodeList = make([]*TreeNode, len(tmpNodes))
		copy(nodeList, tmpNodes)
		tmpNodes = []*TreeNode{}
	}

	return ret
}
