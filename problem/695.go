package problem

/*
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

示例 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

示例 2:

[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。

注意: 给定的矩阵grid 的长度和宽度都不超过 50。

 */

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	for x:=0 ; x<len(grid); x++ {
		for y := 0 ; y < len(grid[0]);y++{
			// 为陆地才遍历
			if grid[x][y] == 1 {
				thisRet := helper(grid, x, y)
				if thisRet > result {
					result = thisRet
				}
			}
		}
	}
	return result
}

func helper(grid [][]int, row, col int) int {
	if row > len(grid) - 1 || row < 0{
		return 0
	}
	if col > len(grid[0]) - 1 || col < 0 {
		return 0
	}
	// left, right, up, down
	if grid[row][col] == 1 {
		grid[row][col] = 0
		return 1 + helper(grid, row, col-1) + helper(grid, row, col+1) + helper(grid, row-1, col) + helper(grid, row+1, col)
	}
	return 0
}

// 连接的不需要归位
//func makeArray(grid [][]int) [][]int{
//	ret := make([][]int, 0)
//	for x:=0; x<len(grid); x++ {
//		thisRow := make([]int, 0)
//		for y:=0; y<len(grid[0]);y++ {
//			thisRow = append(thisRow, grid[x][y])
//		}
//		ret = append(ret, thisRow)
//	}
//	return ret
//}
