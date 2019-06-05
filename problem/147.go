package problem

/*
对链表进行插入排序。


插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。



插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。


示例 1：

输入: 4->2->1->3
输出: 1->2->3->4
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5
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
			if cur.Val < sortNode.Next.Val {
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
