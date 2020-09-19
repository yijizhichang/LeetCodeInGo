package problem

import "sort"

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

*/

/*
排序以后，DFS递归搜索。
在同一个递归深度（即树的层级）迭代当前slice的时候，跳过与之前已经迭代过的相同的值，因为这会导致当前所处的递归层中即将的搜索的路径与之前刚刚搜索过的路径重复。
因为不能重复使用candidates中同一个位置的数字，因此下一层递归时，应当截取当前迭代所处位置的剩余slice。

*/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(nums []int, target int) [][]int {
	ret := [][]int{}
	for i, n := range nums {
		if i > 0 && nums[i-1] == n {
			continue
		} else if target < n {
			break
		} else if target == n {
			ret = append(ret, []int{n})
			continue
		}
		for _, v := range dfs(nums[i+1:], target-n) {
			ret = append(ret, append([]int{n}, v...))
		}
	}
	return ret
}
