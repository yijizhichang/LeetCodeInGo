package array_list

/*
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。



示例 1：

输入：nums = [1,3,5,4,7]
输出：3
解释：最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
示例 2：

输入：nums = [2,2,2,2,2]
输出：1
解释：最长连续递增序列是 [2], 长度为1。

*/
func findLengthOfLCIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	max, cur := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur++
		} else {
			// 断了
			if cur > max {
				max = cur
			}
			cur = 1
		}
	}
	if cur > max {
		max = cur
	}
	return max
}
