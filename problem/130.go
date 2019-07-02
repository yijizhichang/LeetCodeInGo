package problem

/*

给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func solve(board [][]byte)  {
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[i]); j ++ {
			if i != 0 && j != 0 && i != len(board) - 1 && j != len(board[i]) - 1 {
				continue
			}
			if board[i][j] == 'O' {
				mark(board, i, j)
			}
		}
	}
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[i]); j ++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'M' {
				board[i][j] = 'O'
			}
		}
	}

	//for i := 0; i < len(board); i ++ {
	//	for j := 0; j < len(board[i]); j ++ {
	//		fmt.Print(string(board[i][j]))
	//	}
	//	fmt.Println()
	//}

}

func mark(board [][]byte, i int, j int) {
	board[i][j] = 'M'//标记

	//向左
	if j - 1 >= 0 && board[i][j - 1] == 'O' {
		mark(board, i, j - 1)
	}

	//向右
	if j + 1 < len(board[i])  && board[i][j + 1] == 'O' {
		mark(board, i, j + 1)
	}

	//向前
	if i - 1 >= 0  && board[i - 1][j] == 'O' {
		mark(board, i - 1, j)
	}

	//向后
	if i + 1 < len(board)  && board[i + 1][j] == 'O' {
		mark(board, i + 1, j)
	}
}

