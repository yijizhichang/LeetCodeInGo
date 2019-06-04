package problem

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	cur := root
	cur1 := cur
	for cur1!=nil {
		cur = cur1
		if cur.Left != nil {
			cur = cur.Left
			for cur.Right != nil {
				cur = cur.Right
			}
			cur.Right = cur1.Right
			tmp := cur1.Left
			cur1.Left = nil
			cur1.Right = tmp
		}
		cur1 = cur1.Right
	}
}
