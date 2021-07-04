package int

import (
	"strconv"
)

/*
在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 位数字。



注意：n 是正数且在 32 位整数范围内（n < 231）。



示例 1：

输入：3
输出：3
示例 2：

输入：11
输出：0
解释：第 11 位数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是 0 ，它是 10 的一部分。

*/
func findNthDigit(n int) int {
	digits := 1
	base := 9
	for n-digits*base > 0 {
		n -= digits * base
		digits++
		base *= 10
	}
	if digits == 1 {
		return n
	}
	number := 1
	for k := 1; k < digits; k++ {
		number = number * 10
	}
	// 计算出真正的被包含的数
	number += (n - 1) / digits
	idx := (n - 1) % digits
	// 余数就是所处的位数
	numberStr := strconv.Itoa(number)
	bit, _ := strconv.Atoi(string(numberStr[idx]))
	return bit
}
