package double_ptr

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。



示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

*/

func exchange(nums []int) []int {
	x, y := 0, len(nums)-1
	for x < y {
		for x < y && nums[x]%2 != 0 {
			x++
		}
		for x < y && nums[y]%2 == 0 {
			y--
		}
		nums[x], nums[y] = nums[y], nums[x]
	}
	return nums
}
