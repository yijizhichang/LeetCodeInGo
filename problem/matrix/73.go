package matrix

/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

进阶：

一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？

*/

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	// 记录本身第一行，第一列有没有0值
	row0, col0 := false, false

	for _, this := range matrix[0] {
		if this == 0 {
			row0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}

	// 先看非第一行第一列之外的值，有0则把他记录到第一行，第一列里去
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 再从第一行、第一列的记录中来改变行列本来应该变为0的行或列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 再根据row0， col0看是否要把第一行第一列都变为0
	if row0 {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}

	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}

}
