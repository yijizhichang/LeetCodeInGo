package sort

import "sort"

/*
给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。

请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。

如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。

军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。



示例 1：

输入：mat =
[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]],
k = 3

*/

type Point struct {
	point int
	row   int
}

func kWeakestRows(mat [][]int, k int) []int {
	length := len(mat)
	points := make([]Point, length)
	for idx, row := range mat {
		point := 0
		for _, this := range row {
			if this == 0 {
				break
			}
			point++
		}
		points[idx] = Point{point, idx}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].point < points[j].point || points[i].row < points[j].row
	})
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = points[i].row
	}
	return ret
}
