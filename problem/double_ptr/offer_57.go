package double_ptr

/*
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]
示例 2：

输入：nums = [10,26,30,31,47,60], target = 40
输出：[10,30] 或者 [30,10]

*/

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return nil
	}
	left := 0
	right := len(nums) - 1
	//大于目标值的数字全都排除
	for nums[right] > target {
		right--
	}
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		}
		if nums[left] > target-nums[right] {
			right--
		} else {
			left++
		}
	}
	return nil
}
