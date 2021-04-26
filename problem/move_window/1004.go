package move_window
/*
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。

返回仅包含 1 的最长（连续）子数组的长度。

 

示例 1：

输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10

 */

func longestOnes(A []int, K int) int {
	left, right := 0, 0
	maxCount := 0
	zeroCount := 0
	for ;right<len(A);right++ {
		if A[right] == 0 {
			zeroCount ++
		}
		for zeroCount > K {
			if A[left] == 0 {
				zeroCount--
			}
			left++
		}
		maxCount = max(maxCount, right-left+1)
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
