package array_list

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2

*/

func majorityElement(nums []int) int {
	element2Count := make(map[int]int)
	point := len(nums) / 2
	for _, num := range nums {
		if count, ok := element2Count[num]; ok {
			element2Count[num] = count + 1
		} else {
			element2Count[num] = 1
		}
		if element2Count[num] > point {
			return num
		}
	}
	return 0
}

// 摩尔投票算法
func majorityElement(nums []int) int {
	target := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			target = num
		}
		if num == target {
			count++
		} else {
			count--
		}
	}
	return target
}
