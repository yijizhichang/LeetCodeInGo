package problem

/*
你现在手里有一份大小为 N x N 的『地图』（网格） grid，上面的每个『区域』（单元格）都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地，你知道距离陆地区域最远的海洋区域是是哪一个吗？请返回该海洋区域到离它最近的陆地区域的距离。

我们这里说的距离是『曼哈顿距离』（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个区域之间的距离是 |x0 - x1| + |y0 - y1| 。

如果我们的地图上只有陆地或者海洋，请返回 -1。



示例 1：



输入：[[1,0,1],[0,0,0],[1,0,1]]
输出：2
解释：
海洋区域 (1, 1) 和所有陆地区域之间的距离都达到最大，最大距离为 2。
示例 2：



输入：[[1,0,0],[0,0,0],[0,0,0]]
输出：4
解释：
海洋区域 (2, 2) 和所有陆地区域之间的距离都达到最大，最大距离为 4。

*/
func maxDistance(grid [][]int) int {
	var queue []point // all positions of lands
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				queue = append(queue, point{i, j})
			}
		}
	}
	rows := len(grid)
	cols := len(grid[0])
	if len(queue) == 0 || len(queue) == rows*cols {
		return -1
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dist := 0
	for len(queue) > 0 {
		var tempQueue []point
		for _, p := range queue {
			for _, dir := range dirs {
				newX := p.x + dir[0]
				newY := p.y + dir[1]
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols && grid[newX][newY] == 0 {
					tempQueue = append(tempQueue, point{newX, newY})
					grid[newX][newY] = 1
				}
			}
		}
		dist++
		queue = tempQueue
	}
	return dist - 1
}

type point struct {
	x, y int
}
