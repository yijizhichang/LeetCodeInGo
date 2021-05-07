package array_list


/*
给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。

如果不存在满足条件的子数组，则返回 0 。



示例 1：

输入：nums = [8,2,4,7], limit = 4
输出：2
解释：所有子数组如下：
[8] 最大绝对差 |8-8| = 0 <= 4.
[8,2] 最大绝对差 |8-2| = 6 > 4.
[8,2,4] 最大绝对差 |8-2| = 6 > 4.
[8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
[2] 最大绝对差 |2-2| = 0 <= 4.
[2,4] 最大绝对差 |2-4| = 2 <= 4.
[2,4,7] 最大绝对差 |2-7| = 5 > 4.
[4] 最大绝对差 |4-4| = 0 <= 4.
[4,7] 最大绝对差 |4-7| = 3 <= 4.
[7] 最大绝对差 |7-7| = 0 <= 4.
因此，满足题意的最长子数组的长度为 2 。


*/

func longestSubarray(nums []int, limit int) (ans int) {
	var minQ, maxQ []int // 递增的最小队列，递减的最大队列
	left := 0
	for right, v := range nums {
		// 在递增的最小队列中把比当前值大的都移除
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		// 加入当前值
		minQ = append(minQ, v)
		// 在递减的最大队列中把比当前值小的都移除
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		// 尾部加入当前值
		maxQ = append(maxQ, v)

		// maxQ的头值 - minQ的头值即是当前维护窗口中最大绝对值
		// 这里是在维护滑动窗口的合法性
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				// 把最小值过小的剔除
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				// 把最大值过大的剔除
				maxQ = maxQ[1:]
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
