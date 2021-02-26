package matrix

/*
转置矩阵
给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。

矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
 */


func transpose(matrix [][]int) [][]int {
	row, col := len(matrix), len(matrix[0])
	//for i:=0; i < row; i ++ {
	//	for j:=0; j < i && j < col; j ++ {
	//		tmp := matrix[i][j]
	//		matrix[i][j] = matrix[j][i]
	//		matrix[j][i] = tmp
	//	}
	//}
	//return matrix
	ret := make([][]int, col)
	for i:=0; i < col; i ++ {
		ret[i] = make([]int, row)
	}
	for i, vals := range matrix {
		for j, val := range vals {
			ret[j][i] = val
		}
	}
	return ret
}