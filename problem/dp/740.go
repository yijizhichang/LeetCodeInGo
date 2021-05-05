package dp

import "math"

/*
给你一个整数数组nums，你可以对它进行一些操作。

每次操作中，选择任意一个nums[i]，删除它并获得nums[i]的点数。之后，你必须删除每个等于nums[i] - 1或nums[i] + 1的元素。

开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。


示例 1：

输入：nums = [3,4,2]
输出：6
解释：
删除 4 获得 4 个点数，因此 3 也被删除。
之后，删除 2 获得 2 个点数。总共获得 6 个点数。


 */

// 类似打家劫舍
func deleteAndEarn(nums []int) int {
	// 先算出最大值，取数组边界
	maxVal := math.MinInt32
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	dp := make([]int, maxVal+1)
	for _, num := range nums {
		dp[num] += num
	}
	//first, second := dp[0], max(dp[0], dp[1])
	dp[1] = max(dp[0], dp[1])
	for i:=2; i<len(dp); i++ {
		//first, second = second, max(first+dp[i], second)
		dp[i] = max(dp[i-1], dp[i-2]+dp[i])
	}
	return dp[maxVal]
}

