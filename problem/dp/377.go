package dp

/*
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。

题目数据保证答案符合 32 位整数范围。



示例 1：

输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
示例 2：

输入：nums = [9], target = 3
输出：0

*/

/*
可以通过动态规划的方法计算可能的方案数。
用 dp[x] 表示选取的元素之和等于 x 的方案数，目标是求 dp[target]

状态转移方程：
dp[0] = 1 // target为0 什么都不选即可，也是动态规划的边界
对于 dp[i] 如果有 num <= i ，则就有遍历到当前 num 的 dp[i-num]个子集，得出
dp[i] += dp[i-num]
*/

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
