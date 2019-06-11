package problem

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	ret := make([]int, len(nums))
	ret[0] = nums[0]
	max := ret[0]
	for i := 1; i < len(nums); i++ {
		tmp := ret[i-1] + nums[i]
		if tmp > nums[i] {
			ret[i] = tmp
		} else {
			ret[i] = nums[i]
		}
		if max < ret[i] {
			max = ret[i]
		}
	}
	return max
}
