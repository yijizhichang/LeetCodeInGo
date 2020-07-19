package problem

//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//示例:
//输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	mNode := &ListNode{Next: lists[0]}
	for i := 1; i < len(lists); i++ {
		for lists[i] != nil {
			next := lists[i].Next
			insertList(mNode, lists[i])
			lists[i] = next
		}
	}
	return mNode.Next
}

func insertList(list *ListNode, val *ListNode) {
	pre := list
	cur := list.Next
	for pre != nil {
		if cur == nil || val.Val <= cur.Val {
			pre.Next = val
			val.Next = cur
			break
		}
		pre = pre.Next
		cur = cur.Next
	}
}
