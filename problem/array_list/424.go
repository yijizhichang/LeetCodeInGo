package array_list

/*
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换k次。在执行上述操作后，找到包含重复字母的最长子串的长度。

注意：字符串长度 和 k 不会超过104。



示例 1：

输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。
示例 2：

输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。

 */

func characterReplacement(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left, right := 0,0,0
	length := len(s)
	for right < length {
		index := s[right]-'A'
		cnt[index]++
		maxCnt = max(maxCnt, cnt[index])
		if right-left+1 > maxCnt+k {
			cnt[s[left]-'A']--
			left++
		}
		right++
	}
	return right - left + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
