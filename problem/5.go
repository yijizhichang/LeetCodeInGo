package problem

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"
*/

func ishw(i, j int, s string) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindrome(s string) string {
	max := 0
	if len(s) < 2 {
		return s
	}
	ret := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if ishw(i, j, s) {
				if max < j-i {
					max = j - i
					ret = s[i : j+1]
				}
			}
		}
	}
	return ret
}
