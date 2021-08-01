package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("test here!")
}

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
		points[idx] = Point{point*10*k + idx, idx}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].point < points[j].point
	})
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = points[i].row
	}
	return ret
}
