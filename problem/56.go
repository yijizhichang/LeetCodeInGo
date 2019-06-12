package problem

import "sort"

/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2:
输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

type MS [][]int

func (m MS) Len() int {
	return len(m)
}

func (m MS) Less(i, j int) bool {
	if m[i][0] > m[j][0] {
		return false
	}
	return true
}

func (m MS) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func merge(intervals [][]int) [][]int {
	s := len(intervals)
	ret := make([][]int, 0)
	if s == 0 {
		return ret
	}
	sort.Sort(MS(intervals))
	tmp := intervals[0]
	for j := 1; j < s; j++ {
		if tmp[1] >= intervals[j][0] {
			tmp[1] = max(tmp[1], intervals[j][1])
		} else {
			ret = append(ret, tmp)
			tmp = intervals[j]
		}
	}
	ret = append(ret, tmp)
	return ret
}
