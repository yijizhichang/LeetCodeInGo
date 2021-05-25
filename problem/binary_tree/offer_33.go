package binary_tree

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

 

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

 */

func verifyPostorder(postorder []int) bool {
	/*
		思路：
			用递归
			1.找到第一个比当前根结点大的值，取到下标m，以他为左的就是左子树，为右为右子树
			2.确定右子树是不是都小于根结点值，如果不满足，则不是个二叉搜索树
			3.对[i, m-1] 和[m, j-1]重复上述1和2
	*/
	return help(0, len(postorder)-1, postorder)
}

func help(i, j int, postorder []int) bool {
	if i >= j {
		return true
	}
	oriI := i
	for i<j && postorder[i] < postorder[j] {
		i++
	}
	m := i
	for i<j && postorder[i] > postorder[j] {
		i++
	}
	return i == j && help(oriI, m-1, postorder) && help(m, j-1, postorder)
}
