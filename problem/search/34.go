package search

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

*/

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			leftTarget := mid - 1
			rightTarget := mid
			for leftTarget >= 0 && nums[leftTarget] == target {
				leftTarget--
			}
			for rightTarget < len(nums) && nums[rightTarget] == target {
				rightTarget++
			}
			return []int{leftTarget + 1, rightTarget - 1}
		}
	}
	return []int{-1, -1}
}
