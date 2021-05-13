package linked_list

/*
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。

例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。



示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 双指针，先让快指针多走k步
func getKthFromEnd(head *ListNode, k int) *ListNode {
	latter, former := head, head
	for k > 0 {
		former = former.Next
		k--
	}
	for former != nil {
		latter = latter.Next
		former = former.Next
	}
	return latter
}
