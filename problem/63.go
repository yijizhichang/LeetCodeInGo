package problem

/*

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
网格中的障碍物和空位置分别用 1 和 0 来表示。
说明：m 和 n 的值均不超过 100。

示例 1:
输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右

*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if (len(obstacleGrid) == 0) {
		return 0;
	}
	if (obstacleGrid[0][0] == 1) {
		return 0;
	}
	obstacleGrid[0][0] = 1
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = obstacleGrid[i - 1][0]
		}
	}
	for j := 1; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1 {
			obstacleGrid[0][j] = 0
		} else {
			obstacleGrid[0][j] = obstacleGrid[0][j-1]
		}
	}
	for m := 1; m < len(obstacleGrid); m++ {
		for n := 1; n < len(obstacleGrid[0]); n++ {
			if obstacleGrid[m][n] == 1{
				obstacleGrid[m][n] = 0
			}else{
				obstacleGrid[m][n] = obstacleGrid[m-1][n] + obstacleGrid[m][n-1]
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
