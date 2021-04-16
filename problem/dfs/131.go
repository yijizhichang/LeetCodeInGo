package dfs

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

https://leetcode-cn.com/problems/palindrome-partitioning/solution/shou-hua-tu-jie-san-chong-jie-fa-hui-su-q5zjt/
*/

func partition(s string) [][]string {
	res := [][]string{}
	memo := make([][]int, len(s)) // 备忘录：0 代表未计算 1代表是回文 2 代表不是回文
	for i := range memo {
		memo[i] = make([]int, len(s))
	}

	dfs([]string{}, 0, &res, s, memo)
	return res
}
func dfs(temp []string, start int, res *[][]string, s string, memo [][]int) {
	if start == len(s) {
		t := make([]string, len(temp))
		copy(t, temp)
		*res = append(*res, t)
		return
	}

	for i := start; i < len(s); i++ {
		if memo[start][i] == 2 { //不回文，直接跳过
			continue
		}
		if memo[start][i] == 1 || isPali(s, start, i, memo) { //记录为回文 或没有记录但isPali调用为真
			temp = append(temp, s[start:i+1])
			dfs(temp, i+1, res, s, memo)
			temp = temp[:len(temp)-1]
		}
	}
}

func isPali(s string, l, r int, memo [][]int) bool {
	for l < r {
		if s[l] != s[r] {
			memo[l][r] = 2 // 存入memo
			return false
		}
		l++
		r--
	}
	memo[l][r] = 1 // 存入memo
	return true
}
