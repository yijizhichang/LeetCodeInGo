package array_list

import "math"

/*
集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。

给定一个数组 nums 代表了集合 S 发生错误后的结果。

请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

 

示例 1：

输入：nums = [1,2,2,4]
输出：[2,3]
示例 2：

输入：nums = [1,1]
输出：[1,2]

 */

// hash o(n) time, o(n) space
func findErrorNums(nums []int) []int {
	if len(nums) < 2 {
		return nil
	}
	idxs := make([]int, len(nums))
	ret := make([]int, 2)
	for _, num := range nums {
		if idxs[num-1] != 0 {
			// the dup one get
			ret[0] = num
		} else {
			idxs[num-1] = num
		}
	}
	for idx, num := range idxs {
		if num == 0 {
			ret[1] = idx + 1
		}
	}
	return ret
}

// 原地标记
func findErrorNums(nums []int) []int {
	if len(nums) < 2 {
		return nil
	}
	miss, dup := -1, -1
	for _, num := range nums {
		i := int(math.Abs(float64(num))) - 1
		if nums[i] < 0 {
			dup = i+1
		} else {
			nums[i] = -nums[i]
		}
	}
	for idx, num := range nums {
		if num > 0 {
			miss = idx+1
		}
	}
	return []int{dup, miss}
}