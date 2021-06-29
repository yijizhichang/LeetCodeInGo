package array_list

/*
给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。

例如：

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...


示例 1：

输入：columnNumber = 1
输出："A"
示例 2：

输入：columnNumber = 28
输出："AB"

*/

/* 推导出下述公式
number = a0 + a1 + ... + an
number = a0 +26×∑ai(i=1...n−1)×26^(i−1)

a0 := (columnNumber-1)%26 + 1

*/

func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		a0 := (columnNumber-1)%26 + 1
		ans = append(ans, 'A'+byte(a0-1))
		columnNumber = (columnNumber - a0) / 26
	}
	// 由于是从个位输出的，需要反向一下结果
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}
