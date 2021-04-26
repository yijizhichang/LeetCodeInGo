package linked_list

/*
86. 分隔链表
给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。



示例：

输入：head = 1->4->3->2->5->2, x = 3
输出：1->2->2->4->3->5
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	left := &ListNode{
		Val: 0,
	}
	right := &ListNode{
		Val: 0,
	}
	helperL := left
	helperR := right
	for head != nil {
		if head.Val < x {
			helperL.Next = head
			helperL = helperL.Next
		} else {
			helperR.Next = head
			helperR = helperR.Next
		}
		head = head.Next
	}
	helperR.Next = nil
	if right.Next != nil {
		helperL.Next = right.Next
	}
	return left.Next
}
