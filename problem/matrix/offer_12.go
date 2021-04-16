package matrix

/*
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

 

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

 */

func exist(board [][]byte, word string) bool {
	if len(board) <= 0 {
		return false
	}
	m := len(board)
	n := len(board[0])

	mark := make([][]int, 0)
	// init
	for i := 0; i < m; i++ {
		mark = append(mark, make([]int, n))
	}
	directs := make([][]int, 0)
	directs = append(directs, []int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0})
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] == board[i][j] {
				mark[i][j] = 1
				if backtrack(i, j, board, mark, directs, word[1:]) {
					return true
				} else {
					mark[i][j] = 0
				}
			}
		}
	}
	return false
}

func backtrack(i, j int, board [][]byte, mark , directs [][]int, word string) bool {
	if len(word) == 0 {
		return true
	}
	for _, tmp := range directs {
		curI := i + tmp[0]
		curJ := j + tmp[1]
		if curI >= 0 && curI < len(board) && curJ >= 0 && curJ < len(board[0]) && board[curI][curJ] == word[0] {
			// 已经使用的标记跳过
			if mark[curI][curJ] == 1 {
				continue
			}
			// 标记一下当前的
			mark[curI][curJ] = 1
			if backtrack(curI, curJ, board, mark, directs, word[1:]) {
				return true
			} else {
				// 回溯
				mark[curI][curJ] = 0
			}
		}
	}
	return false
}
