package problem

import "math"

/*
	你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
	这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。
	同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

	示例 1:

	输入: [2,3,2]
	输出: 3
	解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

	解题思路：
	由于他是一个闭环，所以可以把问题分解成n[1]~n[n-1] 和 n[2]~n[n] 中取最大值的子问题
	然后分别n[1]~n[n-1]又可以看成线性的打家劫舍1问题
*/

func rob(nums []int) int {
	handler := func(sub []int) int {
		if len(sub) == 0 {
			return 0
		}
		if len(sub) == 1 {
			return sub[0]
		}
		if len(sub) == 2 {
			return int(math.Max(float64(sub[0]), float64(sub[1])))
		}
		dp := make([]int, len(sub))
		dp[0] = sub[0]
		dp[1] = int(math.Max(float64(sub[0]), float64(sub[1])))
		for i := 2; i < len(sub); i++ {
			dp[i] = int(math.Max(float64(dp[i-2]+sub[i]), float64(dp[i-1])))
		}
		return dp[len(dp)-1]
	}
	if len(nums) == 1 {
		return nums[0]
	}
	sub1 := make([]int, 0)
	sub2 := make([]int, 0)
	for idx, cur := range nums {
		if idx <= len(nums)-2 {
			sub1 = append(sub1, cur)
		}
		if idx >= 1 && idx <= len(nums)-1 {
			sub2 = append(sub2, cur)
		}
	}
	return int(math.Max(float64(handler(sub1)), float64(handler(sub2))))
}
