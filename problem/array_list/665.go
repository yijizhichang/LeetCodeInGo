package array_list

/*
给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的： 对于数组中任意的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

 

示例 1:

输入: nums = [4,2,3]
输出: true
解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
示例 2:

输入: nums = [4,2,1]
输出: false
解释: 你不能在只改变一个元素的情况下将其变为非递减数列。

 */

func checkPossibility(nums []int) bool {
	if len(nums) < 2 {
		 return true
	}
	count := 0
	for idx:=0;idx<len(nums)-1;idx++ {
		x, y := nums[idx], nums[idx+1]
		if x > y {
			count++
			if count > 1 {
				 return false
			}
			if idx > 0 && y < nums[idx-1] {
				nums[idx+1] = nums[idx]
			}
		}
	}
	return true
}
