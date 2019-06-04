package problem

func hasCycle(head *ListNode) bool {
	k := head
	m := head
	for k!=nil && m != nil {
		k = k.Next
		if k == nil {
			return false
		}
		k = k.Next
		m = m.Next
		if k == m {
			return true
		}
	}
	return false
}
