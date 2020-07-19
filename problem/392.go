package problem

//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
//你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
//
//字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/is-subsequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	ls := len(s)
	lt := len(t)
	var flag = false
	tmp := 0
	for i := 0; i < ls; i++ {
		for j := tmp; j < lt; j++ {
			if s[i] == t[j] {
				flag = true
				tmp = j + 1
				break
			}
		}
		if !flag || i+1 == ls {
			return flag
		} else {
			flag = false
		}
	}
	return flag
}
