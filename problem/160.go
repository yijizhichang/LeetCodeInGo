package problem

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil {
		return nil
	}
	if headB == nil {
		return nil
	}
	h1,h2 := headA,headB
	has1,has2:= false,false
	for h1 != nil || h2 != nil {
		if h1 == h2 {
			return h1
		}
		h1=h1.Next
		h2=h2.Next
		if h1 == nil && has1 == false {
			h1 = headB
			has1 = true
		}
		if h2 == nil && has2 == false {
			h2 = headA
			has2 = true
		}
	}
	return nil
}
