package problem

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{0,nil}
	tmp := ret
	jw,sum := 0,0
	for (l1 != nil && l2 != nil) || jw != 0 {
		val1,val2 := 0,0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum = val1 + val2 + jw
		jw = sum/10
		tmp.Next = &ListNode{sum%10,nil}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return ret.Next
}
