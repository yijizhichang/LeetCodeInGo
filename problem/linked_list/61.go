package linked_list

/*
61. 旋转链表
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	tmpHead := head
	length := 1
	for tmpHead.Next != nil {
		length++
		tmpHead = tmpHead.Next
	}
	// 连起来成环
	tmpHead.Next = head
	// 从0起下标为n的是新的头节点
	n := length - k%length
	newHead := head
	for i := 0; i < n; i++ {
		if i == n-1 {
			// 断开，使当前节点为最后节点，下一个节点为头节点
			newHead = head.Next
			head.Next = nil
			break
		}
		head = head.Next
	}
	return newHead
}
