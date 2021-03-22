package matrix

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

 */

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return []int{}
	}
	result := make([]int,0)
	upRow,leftCol := 0,0
	downRow := len(matrix)-1
	rightCol := len(matrix[0])-1
	for upRow <= downRow && leftCol <= rightCol {
		// 从左到右
		for i:=leftCol;i<=rightCol;i++{
			result = append(result, matrix[upRow][i])
		}
		upRow ++
		if upRow > downRow {
			break
		}
		// 从上到下
		for i:=upRow;i<=downRow;i++{
			result = append(result, matrix[i][rightCol])
		}
		rightCol--
		if leftCol > rightCol {
			break
		}
		// 从右到左
		for i:=rightCol;i>=leftCol;i--{
			result = append(result, matrix[downRow][i])
		}
		downRow--
		if upRow > downRow {
			break
		}
		// 从下到上
		for i:=downRow;i>=upRow;i--{
			result = append(result, matrix[i][leftCol])
		}
		leftCol++
	}
	return result
}
