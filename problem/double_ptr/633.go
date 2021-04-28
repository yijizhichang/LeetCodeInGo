package double_ptr

import "math"

/*
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。



示例 1：

输入：c = 5
输出：true
解释：1 * 1 + 2 * 2 = 5


双指针
*/

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}
