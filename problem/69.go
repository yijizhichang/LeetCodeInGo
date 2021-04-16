package problem

/*
	实现 int sqrt(int x) 函数。

	计算并返回 x 的平方根，其中 x 是非负整数。

	由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

	示例 1:

	输入: 4
	输出: 2
	示例 2:

	输入: 8
	输出: 2
	说明: 8 的平方根是 2.82842...,
	     由于返回类型是整数，小数部分将被舍去。
*/

// 二分
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	start := 1
	end := x
	for start <= end {
		// 第一次肯定是靠左
		mid := start + (end-start)/2 // 等同start-start/2 + end/2 -> (start+end)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start - 1
}
