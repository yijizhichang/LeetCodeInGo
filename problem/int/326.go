package int

import "math"

/*
给定一个整数，写一个函数来判断它是否是 3的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x

示例 1：

输入：n = 27
输出：true
示例 2：

输入：n = 0
输出：false
示例 3：

输入：n = 9
输出：true
示例 4：

输入：n = 45
输出：false

 */

func isPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

/*
我们还可以使用一种较为取巧的做法。

在题目给定的 32 位有符号整数的范围内，最大的 3 的幂为 3^19 = 1162261467。我们只需要判断 n 是否是 3^{19}的约数即可。

与方法一不同的是，这里需要特殊判断 n 是负数或 0 的情况。

 */

func isPowerOfThree2(n int) bool {
	max := int(math.Pow(3.0, 19.0))
	return n > 0 && max % n == 0
}
