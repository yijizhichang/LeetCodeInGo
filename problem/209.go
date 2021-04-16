package problem

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/

func minSubArrayLen(s int, nums []int) int {
	// 双指针
	if sum(nums) < s {
		return 0
	}
	left, right, ret, sumLr := 0, 0, len(nums), 0
	for right < len(nums) {
		for sumLr < s && right < len(nums) {
			sumLr += nums[right]
			right++
		}
		for sumLr >= s && left >= 0 {
			ret = min(ret, right-left)
			sumLr -= nums[left]
			left++
		}
	}
	return ret
}

func sum(nums []int) int {
	ret := 0
	for _, val := range nums {
		ret += val
	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
