package problem

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

input: 5
1. 1+1+1+1+1
2. 1+1+1+2
3. 1+2+2
4. 2+1+2
5. 2+2+1
6. 1+1+2+1
7. 1+2+1+1
8. 2+1+1+1
 */

// 太多重复的计算
func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 动态规划
func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i:=3;i<=n;i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//func helper(n int, dp []int) int {
//	if dp[n] != 0 {
//		return dp[n]
//	}
//	return helper(n-1,dp) + helper(n-2,dp)
//}

