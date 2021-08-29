package pre_sum

/*
给你一个正整数数组 arr ，请你计算所有可能的奇数长度子数组的和。

子数组 定义为原数组中的一个连续子序列。

请你返回 arr 中 所有奇数长度子数组的和 。

 

示例 1：

输入：arr = [1,4,2,5,3]
输出：58
解释：所有奇数长度子数组和它们的和为：
[1] = 1
[4] = 4
[2] = 2
[5] = 5
[3] = 3
[1,4,2] = 7
[4,2,5] = 11
[2,5,3] = 10
[1,4,2,5,3] = 15
我们将所有值求和得到 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
 */

func sumOddLengthSubarrays(arr []int) int {
	length := len(arr)
	if length < 1 {
		return 0
	}
	if length < 2 {
		return arr[0]
	}
	preSum := make([]int, length)
	preSum[0] = arr[0]
	for i := 1; i < length; i++ {
		preSum[i] = arr[i] + preSum[i-1]
	}

	res := preSum[length-1]
	if length%2 != 0 {
		res += preSum[length-1]
	}
	for i:=3; i<length; i+=2 {
		res += preSum[i-1]
		for left := 0; left<length&&i+left<length; left++ {
			this := preSum[i+left] - preSum[left]
			res += this
		}
	}
	return res
}