package problem

/*
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

//func removeElements(head *ListNode, val int) *ListNode {
//	if head == nil {
//		return head
//	}
//	h := &ListNode{0, head}
//	ret := h
//	for h.Next != nil {
//		if h.Next.Val == val {
//			h.Next = h.Next.Next
//		} else {
//			h = h.Next
//		}
//	}
//	return ret.Next
//}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	helperH := &ListNode{Val: 0, Next: head}
	ret := helperH
	for helperH.Next != nil {
		if helperH.Next.Val == val {
			helperH.Next = helperH.Next.Next
		} else {
			helperH = helperH.Next
		}
	}
	return ret.Next
}
