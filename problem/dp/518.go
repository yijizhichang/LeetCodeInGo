package dp

/*
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。



示例 1:

输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
示例 2:

输入: amount = 3, coins = [2]
输出: 0
解释: 只用面额2的硬币不能凑成总金额3。

*/

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // 总和为0的组合只有不选任何硬币
	for _, coin := range coins {
		// 以当前的硬币为最小值，累加 coin ~ amount 中所有以他最为加量的dp[i-coin]
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
