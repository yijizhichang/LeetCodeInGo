package problem

/*
有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？

如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。

你允许：

装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
示例 1: (From the famous "Die Hard" example)

输入: x = 3, y = 5, z = 4
输出: True
示例 2:

输入: x = 2, y = 6, z = 5
输出: False
*/

// z是x,y最大公约数的倍数，就代表能组成
//func canMeasureWater(x int, y int, z int) bool {
//	if x + y < z {
//		return false
//	}
//	if x == 0 || y == 0{
//		return z == 0 || x + y == z
//	}
//	return z % gcd(x, y) == 0
//}
//
// 最大公约数
//func gcd(a,b int) int {
//	if b == 0 {
//		return a
//	}
//	return gcd(b, a % b)
//}

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}

	if x == 0 || y == 0 {
		return z == (x + y)
	}
	if x > y {
		x, y = y, x
	}

	m := make(map[int]bool)
	m[x] = true
	m[y] = true
	m[x+y] = true

	// y is greater
	greater := y
	for y != x {
		// 不停地从y向x里倒满水
		y = y - x
		if y < 0 {
			y = greater + y
		}
		// y 中剩余的位置量一定满足
		m[y] = true
		// y 剩余 + x全部的量一定满足
		m[y+x] = true
	}

	return m[z]
}
