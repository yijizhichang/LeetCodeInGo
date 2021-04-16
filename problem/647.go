package problem

/*
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。



示例 1：

输入："abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：

输入："aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"

*/

func countSubstrings(s string) int {
	m := len(s)
	res := 0
	for i := 0; i < 2*m-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < m && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}

func countSubstrings(s string) int {
	m := len(s)
	res := 0
	for i := 0; i < m; i++ {
		l, r := i, i
		for l >= 0 && r < m && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	for i := 0; i < m; i++ {
		l, r := i, i+1
		for l >= 0 && r < m && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}
