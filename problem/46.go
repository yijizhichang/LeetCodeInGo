package problem

import "debug/dwarf"

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

*/

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) <= 1 {
		res = append(res, nums)
		return res
	}
	used := make([]int, len(nums))
	dfs(nums, []int{}, used, &res)
	return res
}

func dfs(nums, now, used []int, res *[][]int) {
	if len(now) == len(nums) {
		tmp := make([]int, 0)
		tmp = append(tmp, now...)
		*res = append(*res, tmp)
		return
	}
	for idx, val := range nums {
		if used[idx] == 0 {
			used[idx] = 1
			now = append(now, val)
			dfs(nums, now, used, res)
			used[idx] = 0
			now = now[:len(now)-1]
		}
	}
}
