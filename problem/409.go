package problem

/*
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

 */

func longestPalindrome(s string) int {
	letter2count := make(map[rune]int)
	for _, letter := range s {
		if _, ok := letter2count[letter];ok {
			letter2count[letter] += 1
		} else {
			letter2count[letter] = 1
		}
	}
	ret := 0
	hasOdd := false
	for _, count := range letter2count {
		if count % 2 == 0 {
			// 可以整除说明可以成对子
			ret += count
		} else if count > 1{
			// 大于1但是单数
			ret += count / 2 * 2
			hasOdd = true
		} else {
			hasOdd = true
		}
	}
	if hasOdd {
		ret += 1
	}
	return ret
}
