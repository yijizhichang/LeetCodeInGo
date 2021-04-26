package move_window

import "math"

/*
给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。



示例 1：

输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
示例 2：

输入：s = "ababbc", k = 2
输出：5
解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

*/
func longestSubstring(s string, k int) int {
	return recursionHelper(0, len(s), k, s)
}

func recursionHelper(start, end, k int, s string) int {
	if end-start < k {
		return 0
	}
	count := make([]int, 26)
	for i := start; i < end; i++ {
		count[s[i]-'a']++
	}
	for pos, thisTime := range count {
		if thisTime > 0 && thisTime < k {
			for j := start; j < end; j++ {
				if int(s[j]-'a') == pos {
					left := recursionHelper(start, j, k, s)
					right := recursionHelper(j+1, end, k, s)
					return max(left, right)
				}
			}
		}
	}
	return end - start
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
