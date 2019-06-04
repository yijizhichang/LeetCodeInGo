package problem

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	ret := make([]int,0)
	h := head
	for h!= nil {
		ret = append(ret,h.Val)
		h=h.Next
	}
	l,r := 0,len(ret)-1
	for l<r {
		if ret[l] != ret[r] {
			return false
		}
		l++
		r--
	}
	return true
}
