package array_list

/*
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

示例 1:

输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。

*/

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	if length < 1 {
		return []int{}
	}
	if length == 1 {
		return []int{-1}
	}
	ret := make([]int, 0)
	for idx, now := range nums {
		j := idx + 1
		j %= length
		for {
			if nums[j] > now {
				ret = append(ret, nums[j])
				break
			}
			j++
			j = j % length
			if j == idx {
				ret = append(ret, -1)
				break
			}
		}
	}
	return ret
}

// 单调栈
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{}
	for i := 0; i < n*2-1; i++ {
		// 当前值如果比栈顶的大，依次取出，直到栈顶比当前元素大为止
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		// 将当前值的下标存入栈顶
		stack = append(stack, i%n)
	}
	return ans
}
