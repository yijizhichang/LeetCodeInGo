package sort

/*
数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。



示例 1：

输入：[1,2,5,9,5,9,5,5,5]
输出：5
示例 2：

输入：[3,2]
输出：-1


在 Boyer-Moore投票算法中，
遇到相同的数则将 count加 1，遇到不同的数则将 count 减 1。根据主要元素的定义，
主要元素的出现次数大于其他元素的出现次数之和，因此在遍历过程中，
主要元素和其他元素两两抵消，最后一定剩下至少一个主要元素，此时 candidate为主要元素，且 count≥1。


*/
func majorityElement(nums []int) int {
	count, target := 0, -1
	for _, num := range nums {
		if count == 0 {
			count++
			target = num
		} else {
			if target == num {
				count++
			} else {
				count--
			}
		}
	}
	if count <= 0 {
		return -1
	}
	count = 0
	// 因为不一定存在主要元素，判定target是否为主要元素
	for _, num := range nums {
		if num == target {
			count++
		}
	}
	if count > len(nums)/2 {
		return target
	} else {
		return -1
	}
}
