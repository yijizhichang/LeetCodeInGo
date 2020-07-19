package problem

/*
在一个 8 x 8 的棋盘上，有一个白色车（rook）。也可能有空方块，白色的象（bishop）和黑色的卒（pawn）。它们分别以字符 “R”，“.”，“B” 和 “p” 给出。大写字符表示白棋，小写字符表示黑棋。

车按国际象棋中的规则移动：它选择四个基本方向中的一个（北，东，西和南），然后朝那个方向移动，直到它选择停止、到达棋盘的边缘或移动到同一方格来捕获该方格上颜色相反的卒。另外，车不能与其他友方（白色）象进入同一个方格。

返回车能够在一次移动中捕获到的卒的数量

*/

func numRookCaptures(board [][]byte) int {
	if len(board) < 1 {
		return 0
	}
	row := len(board)
	col := len(board[0])
	targetR := 0
	targetC := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'R' {
				targetR = i
				targetC = j
			}
		}
	}
	ret := 0
	row = targetR
	for row > 0 {

		if board[row][targetC] == 'B' {
			break
		} else if board[row][targetC] == 'p' {
			ret += 1
			break
		}
		row -= 1
	}
	col = targetC
	for col > 0 {

		if board[targetR][col] == 'B' {
			break
		} else if board[targetR][col] == 'p' {
			ret += 1
			break
		}
		col -= 1
	}
	row = targetR
	for row < len(board) {

		if board[row][targetC] == 'B' {
			break
		} else if board[row][targetC] == 'p' {
			ret += 1
			break
		}
		row += 1
	}
	col = targetC
	for col < len(board[0]) {

		if board[targetR][col] == 'B' {
			break
		} else if board[targetR][col] == 'p' {
			ret += 1
			break
		}
		col += 1
	}
	return ret
}
