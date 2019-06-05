package problem

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
 */

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	ret := make([]int, 0)
	h := head
	for h != nil {
		ret = append(ret, h.Val)
		h = h.Next
	}
	l, r := 0, len(ret)-1
	for l < r {
		if ret[l] != ret[r] {
			return false
		}
		l++
		r--
	}
	return true
}
