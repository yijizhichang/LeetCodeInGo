package linked_list

/*
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
 

示例 1：


输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	preNode := head
	headNode := preNode.Next
	for headNode != nil {
		if headNode.Val == val {
			preNode.Next = headNode.Next
		} else {
			preNode = preNode.Next
		}
		headNode = headNode.Next
	}
	return head
}
