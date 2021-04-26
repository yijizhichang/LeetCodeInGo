package binary_tree

/*
给定一个二叉树，原地将它展开为一个单链表。



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
	flatten(root.Left)
	flatten(root.Right)
	/**** 后序遍历位置 ****/
	// 1、左右子树已经被拉平成一条链表
	left := root.Left
	right := root.Right

	// 2、将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 3、将原先的右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
