package move_window

/*
给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

 

示例 1：

输入：[1, 2, 2, 3, 1]
输出：2
解释：
输入数组的度是2，因为元素1和2的出现频数最大，均为2.
连续子数组里面拥有相同度的有如下所示:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组[2, 2]的长度为2，所以返回2.

 */

type Entity struct {
	Count int
	L int
	R int
}
func findShortestSubArray(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	count := make(map[int]Entity, 0)
	for idx, this := range nums {
		if e, ok := count[this]; ok {
			e.Count++
			e.R = idx
			count[this] = e
		} else {
			count[this] = Entity{
				Count: 1,
				L: idx,
				R: idx,
			}
		}
	}
	maxCount, ans := 0, 1
	for _, v := range count {
		if v.Count > maxCount {
			ans = v.R - v.L + 1
			maxCount = v.Count
		} else if v.Count == maxCount {
			ans = min(ans, v.R - v.L+1)
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}