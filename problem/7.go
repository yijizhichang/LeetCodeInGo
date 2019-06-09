package problem

import "math"

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:
输入: 123
输出: 321

示例 2:
输入: -123
输出: -321

示例 3:
输入: 120
输出: 21
*/

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}
