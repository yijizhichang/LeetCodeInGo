package array_list

/*
给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。

请返回 nums 的动态和。

 

示例 1：

输入：nums = [1,2,3,4]
输出：[1,3,6,10]
解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。

 */

// 前缀和
func runningSum(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	for i:=1; i <= length-1;i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
