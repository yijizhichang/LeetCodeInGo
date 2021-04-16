package problem

/*
给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters 相同



示例 1：

输入：s = "bcabc"
输出："abc"
示例 2：

输入：s = "cbacdcbc"
输出："acdb"

*/

func removeDuplicateLetters(s string) string {
	valM := make([]int32, 26)
	stack := make([]int32, 0)
	for _, this := range s {
		valM[this-'a']++
	}
	for _, this := range s {
		for j := len(stack) - 1; j > 0; j-- {
			if stack[j] > this {
				// 当前字符序列要小于栈顶字符序列，所以需要向i后面查找是否还有和栈顶一样的值
				// 如果i后面还有栈顶相同元素，则移除栈顶；否则不移除，加入当前元素即可
				if valM[stack[j]] > 1 {
					stack = stack[:j]
					valM[stack[j]]--
				} else {
					// 写入当前元素并推出栈顶循环
					stack = append(stack, this)
					break
				}
			}
		}
	}
	ret := ""
	for _, this := range stack {
		ret += string(this + 'a')
	}
	return ret
}
