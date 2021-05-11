package xor

/*
给你一个整数数组 perm ，它是前 n 个正整数的排列，且 n 是个 奇数 。

它被加密成另一个长度为 n - 1 的整数数组 encoded ，满足 encoded[i] = perm[i] XOR perm[i + 1] 。
比方说，如果 perm = [1,3,2] ，那么 encoded = [2,1] 。

给你 encoded 数组，请你返回原始数组 perm 。题目保证答案存在且唯一。



示例 1：

输入：encoded = [3,1]
输出：[1,2,3]
解释：如果 perm = [1,2,3] ，那么 encoded = [1 XOR 2,2 XOR 3] = [3,1]
示例 2：

输入：encoded = [6,5,4,6]
输出：[2,4,1,5,3]

*/

func decode(encoded []int) []int {
	// 由于prem是前n个正整数的序列，所以prem所有元素的异或结果是可以得出的
	n := len(encoded)
	total := 0
	for i := 1; i <= n+1; i++ {
		total ^= i
	}
	// 推公式可得，encoded中奇数异或即是prem中除了第一个元素之外所有的元素异或结果
	oddSum := 0
	for i := 1; i < n; i += 2 {
		oddSum ^= encoded[i]
	}
	perm := make([]int, n+1)
	// 现在可得首个元素
	perm[0] = total ^ oddSum
	for i, v := range encoded {
		perm[i+1] = perm[i] ^ v
	}
	return perm

}
