package array_list

import (
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

*/

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for _, two := range twoSum(nums, i+1, -nums[i]) {
			ret = append(ret, append(two, nums[i]))
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}

func twoSum(nums []int, idx int, target int) [][]int {
	i := idx
	j := len(nums) - 1
	ret := make([][]int, 0)
	for i < j {
		left := nums[i]
		right := nums[j]
		sum := left + right
		if sum > target {
			for i < j && right == nums[j] {
				j--
			}
		} else if sum < target {
			for i < j && left == nums[i] {
				i++
			}
		} else {
			ret = append(ret, []int{left, right})
			for i < j && right == nums[j] {
				j--
			}
			for i < j && left == nums[i] {
				i++
			}
		}
	}
	return ret
}
