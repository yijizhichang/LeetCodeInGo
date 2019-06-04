package problem

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head == nil {
		return head
	}
	h := &ListNode{0,head}
	ret := h
	for h.Next != nil {
		if h.Next.Val == val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
	return ret.Next
}
