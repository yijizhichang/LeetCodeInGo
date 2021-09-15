package search

/*
峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

你可以假设nums[-1] = nums[n] = -∞ 。

你必须实现时间复杂度为 O(log n) 的算法来解决此问题。



示例 1：

输入：nums = [1,2,3,1]
输出：2
解释：3 是峰值元素，你的函数应该返回其索引 2。
示例2：

输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5
解释：你的函数可以返回索引 1，其峰值元素为 2； 或者返回索引 5， 其峰值元素为 6。

1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
对于所有有效的 i 都有 nums[i] != nums[i + 1]

 */


func findPeakElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}
	if length == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right - left)/2 + left
		i, j := mid - 1, mid + 1
		if i < 0 || j > length -1 {
			return mid
		}
		// 用中间值和左右边的值比较，如果比左右都大，那就是峰值
		// 如果左>中>右： 那么right左移， 因为目前已经保证了左>中，左有机会成为峰值
		// 反之左<中<右： 那么left右移， 因为目前已经保证了中<右，右有机会成为峰值
		if nums[mid] > nums[j] && nums[mid] > nums[i] {
			return mid
		} else if nums[mid] > nums[j] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return 0
}