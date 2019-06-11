package problem

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	originNode := &ListNode{Next: head}
	//指向当前位置的指针
	preCur := head
	cur := head.Next
	for cur != nil {
		sortNode := originNode
		for sortNode.Next != cur {
			if cur.Val < sortNode.Next.Val{
				preCur.Next = cur.Next
				cur.Next = sortNode.Next
				sortNode.Next = cur
				// 如果有节点插入
				cur = preCur.Next
				break
			}
			sortNode = sortNode.Next
		}
		if sortNode.Next == cur {
			// 如果没有节点插入
			preCur = preCur.Next
			cur = cur.Next
		}
	}
	return originNode.Next
}
