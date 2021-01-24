package string

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."

*/

func replaceSpace(s string) string {
	var str string
	start := 0
	for key, value := range s {
		if value == ' ' {
			str += s[start:key]
			str += "%20"
			start = key + 1
		}
	}
	str += s[start:]
	return str
}
