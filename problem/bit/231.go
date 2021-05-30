package bit

/*
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。

如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。

*/

func isPowerOfTwo(n int) bool {
	// 2 的 幂次方按位排列只有一个1 only one 1
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}
