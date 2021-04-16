package problem

/*
在给定的网格中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。

输入：[[2,1,1],[1,1,0],[0,1,1]]
输出：4
*/

func orangesRotting(grid [][]int) int {
	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	row := len(grid)
	col := len(grid[0])
	ret := 0
	freshNum := 0
	queue := make([]point, 0)
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 2 {
				queue = append(queue, point{x: i, y: j})
			} else if grid[i][j] == 1 {
				freshNum++
			}
		}
	}
	for len(queue) > 0 {
		newQueue := make([]point, 0)
		for _, bad := range queue {
			for _, this := range dir {
				nextX := bad.x + this[0]
				nextY := bad.y + this[1]
				if nextX >= 0 && nextX < row && nextY >= 0 && nextY < col && grid[nextX][nextY] == 1 {
					grid[nextX][nextY] = 2
					newQueue = append(newQueue, point{x: nextX, y: nextY})
					freshNum--
				}
			}
		}
		queue = newQueue
		ret++
	}
	if freshNum > 0 {
		return -1
	}
	return ret
}
