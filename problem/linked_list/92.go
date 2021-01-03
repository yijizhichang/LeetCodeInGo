package linked_list

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

var nexter *ListNode = nil

// 反转前n个
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 到最后一个了，下一个不反转了
		nexter = head.Next
		return head
	}
	// 以 head.next 为起点，需要反转前 n - 1 个节点
	last := reverseN(head.Next, n-1)
	// 当前head.Next的下一个节点反向
	head.Next.Next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next = nexter
	return last
}
