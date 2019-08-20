package problem

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

*/
func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return head
    }
    
    root := &ListNode{Next: head}
    small := root
    pre := root
    curr := head
    for curr != nil {
        if curr.Val < x {
            pre.Next = curr.Next
            tmp := small.Next
            small.Next = curr
            curr.Next = tmp
            small = small.Next
        }
        pre = curr
        curr = curr.Next
    }
    return root.Next
}
