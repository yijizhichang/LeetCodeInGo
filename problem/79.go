package problem

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.

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
