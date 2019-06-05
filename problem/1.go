package problem

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
 */

/*
	O(n2)
 */
func twoSum1(nums []int, target int) []int {
	ret := make([]int, 0)
	if len(nums) < 2 {
		return ret
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ret = append(ret, i)
				ret = append(ret, j)
			}
		}
	}
	return ret
}

/*
  O(n)
 */
func twoSum2(nums []int, target int) []int {
	ret := make([]int, 0)
	mp := make(map[int]int)
	if len(nums) < 2 {
		return ret
	}
	for i := 0; i < len(nums); i++ {
		if v, ok := mp[target-nums[i]]; ok {
			ret = append(ret, i, v)
		} else {
			mp[nums[i]] = i
		}
	}
	return ret
}
