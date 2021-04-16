package array_list

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for _, three := range threeSumHelper(nums, i+1, target-nums[i]) {
			ret = append(ret, append(three, nums[i]))
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}

func threeSumHelper(nums []int, idx, target int) [][]int {
	ret := make([][]int, 0)
	for i := idx; i < len(nums); i++ {
		for _, two := range twoSum(nums, i+1, target-nums[i]) {
			ret = append(ret, append(two, nums[i]))
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}
