package array_list
/*
给定一个二进制数组， 计算其中最大连续1的个数。

示例 1:

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
注意：

输入的数组只包含 0 和1。
输入数组的长度是正整数，且不超过 10,000。

 */

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	ret := 0
	for right < len(nums) {
		if nums[left] != 1 {
			right++
			left++
			continue
		}
		if nums[left] == nums[right] {
			right++
			ret = max(ret, right-left)
		} else {
			right++
			left = right
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
