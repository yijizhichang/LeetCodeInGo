package problem

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := head
	length := 1
	for head.Next != nil {
		length++
		head = head.Next
	}
	head.Next = newHead
	newPos := length - k%length
	for i := 1; i < newPos; i++ {
		newHead = newHead.Next
	}
	res := newHead.Next
	newHead.Next = nil
	return res

}

