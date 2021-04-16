package sort

/*
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4

*/

func findKthLargest(nums []int, k int) int {
	left := 0
	n := len(nums)
	right := n - 1
	for {
		idx := partition(nums, left, right)
		// 从小到大第idx小的就是第n-k个
		if idx == n-k {
			return nums[idx]
		} else if idx < n-k {
			left = idx + 1
		} else {
			right = idx - 1
		}
	}
}

func partition(nums []int, left, right int) int {
	sentinal := nums[left]
	l := left
	for l != right {
		for l < right && nums[right] >= sentinal {
			right--
		}
		for l < right && nums[l] <= sentinal {
			l++
		}

		if l < right {
			nums[l], nums[right] = nums[right], nums[l]
		}
	}
	nums[left], nums[l] = nums[l], nums[left]
	return l
}
