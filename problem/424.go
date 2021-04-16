package problem

import "math"

/*
	题目描述：
	给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

	注意:
	字符串长度 和 k 不会超过 104。

	示例 1:
	输入:
	s = "ABAB", k = 2

	输出:
	4

	解释:
	用两个'A'替换为两个'B',反之亦然。
*/

// 思路如下：
// 这里采用滑动窗口，当滑动窗口达到一个最大值后，窗口一直前进
// 当窗口内记载的出现次数最大字符大于上一次的出现最大次数字符时，窗口长度增加
// 否则，窗口保持原来长度前进一个单位

func characterReplacement(s string, k int) int {
	maxCount := 0
	l := 0
	var r int
	count := make([]int, 26)
	for r = 0; r < len(s); r++ {
		count[s[r]-'A']++
		maxCount = int(math.Max(float64(maxCount), float64(count[s[r]-'A'])))
		if r-l+1-maxCount > k {
			count[s[l]-'A']--
			l++
		}
	}
	return r - l
}
