package problem

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/

func hmax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func hmin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0
	for l < r {
		max = hmax(max, hmin(height[r], height[l])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
