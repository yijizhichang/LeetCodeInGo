package problem

/*
在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。

请你返回最终形体的表面积。



示例 1：

输入：[[2]]
输出：10
示例 2：

输入：[[1,2],[3,4]]
输出：34
示例 3：

输入：[[1,0],[0,2]]
输出：16
示例 4：

输入：[[1,1,1],[1,0,1],[1,1,1]]
输出：32
示例 5：

输入：[[2,2,2],[2,1,2],[2,2,2]]
输出：46


提示：

1 <= N <= 50
0 <= grid[i][j] <= 50

*/

func surfaceArea(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	// 正方形个数
	block := 0
	// 被盖住的面(这里只记录了一边的，其实是结果要*2)
	cover := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			block += grid[i][j]
			// 当前坐标重叠的面
			if grid[i][j] > 1 {
				cover += grid[i][j] - 1
			}
			// 与这行上一列重叠的面
			if j > 0 {
				cover += min(grid[i][j-1], grid[i][j])
			}
			// 与上行这列重叠的面
			if i > 0 {
				cover += min(grid[i-1][j], grid[i][j])
			}
		}
	}
	return block*6 - cover*2
}
