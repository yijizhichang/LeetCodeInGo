package problem

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

*/

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s) < 1 {
		return false
	}
	res := make([]string, 0)
	var left2right = map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	for _, this := range s {
		strThis := string(this)
		if len(res) == 0 {
			res = append(res, strThis)
		} else if _, ok := left2right[strThis]; ok {
			res = append(res, strThis)
		} else {
			if left2right[res[len(res)-1]] != strThis {
				return false
			} else {
				res = res[:len(res)-1]
			}
		}
	}
	return len(res) == 0
}
