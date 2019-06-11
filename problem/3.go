package problem

/*
给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	max := 0
	for i := 0; i < n; i++ {
		mp := make(map[uint8]int)
		mp[s[i]] = 1
		for j := i + 1; j < n; j++ {
			if _, ok := mp[s[j]]; !ok {
				mp[s[j]] = 1
			} else {
				break
			}
		}
		if max < len(mp) {
			max = len(mp)
		}
	}
	return max
}
