package problem

import "container/list"

/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口 k 内的数字。滑动窗口每次只向右移动一位。
返回滑动窗口最大值。

示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
 */

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 1 {
		return nums
	}

	//list是一个双向链表，存储当前窗口的最大值。

	queue := list.New()
	//result 存放所有滑动窗口的最大值
	result := make([]int, 0, len(nums)-k+1)

	for i, v := range nums {
		for queue.Len() > 0 {
			// 如果当前数字大于队列尾，则删除队列尾，直到当前数字小于等于队列尾，或者队列空
			if x := queue.Back(); nums[x.Value.(int)] <= v {
				queue.Remove(x)
			} else {
				break
			}
		}
		//当前元素入队列
		queue.PushBack(i)
		// 如果队列头元素不在滑动窗口中了，就删除头元素
		for x, j := queue.Front(), i-k+1; x.Value.(int) < j; x = queue.Front() {
			queue.Remove(x)
		}
		//形成第一个滑动窗口时，才把队首元素放入结果集
		if i >= k-1 {
			result = append(result, nums[queue.Front().Value.(int)])
		}
	}

	return result
}
