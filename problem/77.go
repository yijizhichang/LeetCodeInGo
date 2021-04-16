package problem

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

*/

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	var dfs func(cur int)
	dfs = func(cur int) {
		if len(tmp)+n-cur+1 < k {
			// 剪枝
			return
		}
		if len(tmp) == k {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}
		// 选取当前数字
		tmp = append(tmp, cur)
		dfs(cur + 1)
		tmp = tmp[:len(tmp)-1]
		// 不选当前数字
		dfs(cur + 1)
	}
	dfs(1)
	return res
}
