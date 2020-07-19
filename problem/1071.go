package problem

/*
对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。

返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。



示例 1：

输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"
示例 2：

输入：str1 = "ABABAB", str2 = "ABAB"
输出："AB"
示例 3：

输入：str1 = "LEET", str2 = "CODE"
输出：""


提示：

1 <= str1.length <= 1000
1 <= str2.length <= 1000
str1[i] 和 str2[i] 为大写英文字母

*/

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	var remainder = modOfStrings(str1, str2)
	if remainder == "" {
		return str2
	} else if remainder == str1 {
		return ""
	}
	return gcdOfStrings(str2, remainder)
}
func modOfStrings(str1, str2 string) string {
	length := len(str2)
	var remainder = str1
	for {
		if len(remainder) < length || remainder[:length] != str2 {
			break
		}
		remainder = remainder[length:]
	}
	return remainder
}
