package problem

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0
*/

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for i < len(nums) {
		if nums[i] < target {
			i++
		} else if nums[i] == target {
			return i
		} else {
			break
		}
	}
	return i
}

// 二分法
func searchInsert2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	mid := -1
	for left <= right && right >= 0 && left <= len(nums) {
		mid = left + (right-left)/2 // 防止int溢出
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
			mid = left
		} else {
			return mid
		}
	}
	return mid
}
