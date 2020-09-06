package problem

/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

*/

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	now := make([]*TreeNode, 0)
	now = append(now, root)
	floorVal := make([]int, 0)
	for len(now) > 0 {
		tmp := make([]*TreeNode, 0)
		// 取出当前行的值
		tmp = now
		now = []*TreeNode{}
		for len(tmp) > 0 {
			thisFloor := tmp[0]
			tmp = tmp[1:]
			floorVal = append(floorVal, thisFloor.Val)
			if thisFloor.Left != nil {
				now = append(now, thisFloor.Left)
			}
			if thisFloor.Right != nil {
				now = append(now, thisFloor.Right)
			}
		}
		// 当前行结束
		res = append(res, floorVal)
		floorVal = []int{}
	}
	//l := len(res)
	//reverseRes := make([][]int, l)
	//for idx, val := range res {
	//	reverseRes[l-idx-1] = val
	//}
	//return reverseRes
	reverseSli(res)
	return res
}

func reverseSli(input [][]int) {
	l := len(input)
	if l < 2 {
		return
	}
	left, right := 0, 0
	if l%2 == 0 {
		// even
		left, right = l/2-1, l/2

	} else {
		// odd
		left, right = l/2, l/2
	}
	for left >= 0 && right < l {
		input[left], input[right] = input[right], input[left]
		left--
		right++
	}
}
