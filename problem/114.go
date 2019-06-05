package problem

/*
给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
 */

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	cur1 := cur
	for cur1 != nil {
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
