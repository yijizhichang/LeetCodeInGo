package linked_list

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

图示两个链表在节点 c1 开始相交：



题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。

 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	tmpA := headA
	tmpB := headB
	for tmpA != nil || tmpB != nil {
		if tmpA == nil {
			tmpA = headB
		}
		if tmpB == nil {
			tmpB = headA
		}
		if tmpA == tmpB {
			return tmpA
		}
		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}
	return nil
}