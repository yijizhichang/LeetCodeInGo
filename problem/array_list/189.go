package array_list

/*
189. 旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。



进阶：

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？


示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
*/
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

// 环状替换
func rotate(nums []int, k int) {
	length := len(nums)
	for start, count := 0, 0; count < length; start++ {
		curr := start
		prev := nums[start]
		for {
			next := (curr + k) % length
			nextTmp := nums[next]
			nums[next] = prev // 换位置
			prev = nextTmp
			curr = next // 保持被替换的元素
			count++
			// 如果指向当前轮次的首个下标，则跳出当论，下标++
			if start == curr {
				break
			}
		}
	}
}
