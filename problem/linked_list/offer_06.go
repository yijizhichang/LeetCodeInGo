package linked_list

/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	val := make([]int, 0)
	for head != nil {
		val = append(val, head.Val)
		head = head.Next
	}
	length := len(val)
	for i := 0; i < len(val)/2; i++ {
		val[i], val[length-i-1] = val[length-i-1], val[i]
	}
	return val
}
