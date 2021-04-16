package problem

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

*/

func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	dp := make([][]int, row)
	tmp := make([]int, col)
	for i := 0; i < row; i++ {
		dp[i] = append(dp[i], tmp...)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 || j == 0 {
				if i == 0 {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				} else {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				}
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[row-1][col-1]
}
