package binary_tree

/*
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]

*/

func levelOrder(root *TreeNode) []int {
	ret := make([]int, 0)
	thisFloor := make([]*TreeNode, 0)
	if root == nil {
		return ret
	}
	thisFloor = append(thisFloor, root)
	for len(thisFloor) > 0 {
		this := thisFloor[0]
		ret = append(ret, this.Val)
		if this.Left != nil {
			thisFloor = append(thisFloor, this.Left)
		}
		if this.Right != nil {
			thisFloor = append(thisFloor, this.Right)
		}
		thisFloor = thisFloor[1:]
	}
	return ret
}


