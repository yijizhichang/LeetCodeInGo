package binary_tree

/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

 

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]

 */


func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	thisFloor := []*TreeNode{root}
	flag := true
	for len(thisFloor) > 0 {
		length := len(thisFloor)
		thisFloorVal := make([]int, length)
		for idx:=0; idx<length; idx++ {
			this := thisFloor[idx]
			if flag {
				thisFloorVal[idx] = this.Val
			} else {
				thisFloorVal[length-idx-1] = this.Val
			}

			if this.Left != nil {
				thisFloor = append(thisFloor, this.Left)
			}
			if this.Right != nil {
				thisFloor = append(thisFloor, this.Right)
			}
		}
		flag = !flag
		thisFloor = thisFloor[length:]
		ret = append(ret, thisFloorVal)
	}
	return ret
}