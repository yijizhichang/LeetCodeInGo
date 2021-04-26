package array_list

/*
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，
但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

*/

// 原地置换
func findRepeatNumber(nums []int) int {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			tmp = nums[i]
			nums[i] = nums[tmp]
			nums[tmp] = tmp
		}
	}
	return -1
}

func findRepeatNumber(nums []int) int {
	m := make(map[int]int, 0)
	for _, this := range nums {
		if _, ok := m[this]; ok {
			return this
		} else {
			m[this] = 1
		}
	}
	return -1
}
