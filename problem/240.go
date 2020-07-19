package problem

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。

*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i := len(matrix) - 1
	j := 0
	// 从左下角开始走
	for j < len(matrix[0]) && i >= 0 {
		if matrix[i][j] > target {
			// 向上走
			i--
		} else if matrix[i][j] < target {
			// 向右走
			j++
		} else {
			return true
		}
	}
	return false
}
