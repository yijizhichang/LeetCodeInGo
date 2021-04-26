package dfs

import "fmt"

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

 

示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1

 */

func movingCount(m int, n int, k int) int {
	// 使用匿名函数 - 需要提前声明才能调用自身
	var dfs func(x, y int) (res int)

	// 用Map存储已访问过的坐标 - 利用Map的特性
	visited := make(map[string]bool)

	// 计算数位和
	digitSum := func(x int) (res int) {
		for ; x > 0; x = x / 10 {
			res += x % 10
		}
		return
	}

	// 深度优先搜索
	dfs = func(x, y int) (res int) {
		// 判断坐标是否越界
		if y < 0 || y > n-1 {
			return
		}
		if x < 0 || x > m-1 {
			return
		}

		// 判断坐标是否已经访问过
		if _, isExist := visited[fmt.Sprintf("%d,%d", x, y)]; isExist {
			return
		}

		// 判断数位和是否大于k
		if (digitSum(y) + digitSum(x)) > k {
			return
		}

		// 将该坐标加入已访问
		visited[fmt.Sprintf("%d,%d", x, y)] = true
		res += 1

		// 搜索相邻的坐标
		res += dfs(x+1, y)
		res += dfs(x, y+1)
		return
	}

	// 执行DFS算法
	return dfs(0, 0)
}
