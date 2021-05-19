package xor

import "sort"

/*
给你一个二维矩阵 matrix 和一个整数 k ，矩阵大小为 m x n 由非负整数组成。

矩阵中坐标 (a, b) 的 值 可由对所有满足 0 <= i <= a < m 且 0 <= j <= b < n 的元素 matrix[i][j]（下标从 0 开始计数）执行异或运算得到。

请你找出 matrix 的所有坐标中第 k 大的值（k 的值从 1 开始计数）。

ex1:
输入：matrix = [[5,2],[1,6]], k = 1
输出：7
解释：坐标 (0,1) 的值是 5 XOR 2 = 7 ，为最大的值。
 */

func kthLargestValue(matrix [][]int, k int) int {
	row, col := len(matrix), len(matrix[0])
	pre := make([][]int, row+1)
	pre[0] = make([]int, col+1)
	result := make([]int, 0, row*col)
	for i:=0; i<row; i++ {
		pre[i+1] = make([]int, col+1)
		for j:=0; j<col; j++ {
			pre[i+1][j+1] = pre[i][j]^pre[i+1][j]^pre[i][j+1]^matrix[i][j]
			result = append(result, pre[i+1][j+1])
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i]<result[j]
	})
	//rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	//targetPos := len(result) - k
	//left, right := 0, len(result)-1
	//for left <= right {
	//	pos := Partition(result, left, right)
	//	if pos < targetPos {
	//		left = pos+1
	//	} else if pos > targetPos {
	//		right = pos-1
	//	} else {
	//		return result[pos]
	//	}
	//}
	return result[len(result)-k]
}

func Partition(tar []int, low, high int) int {
	left, right := low, high
	pivot := tar[left]
	// 终止条件是左右指向同一个值
	for left != right {
		// 一定先找右边的
		// 右边开始，找到第一个小于哨兵值的点
		for left < right && pivot <= tar[right] {
			right--
		}
		// 从左边开始，找到第一个大于哨兵值的点
		for left < right && pivot >= tar[left] {
			left++
		}
		if left < right {
			// 交换
			tar[left], tar[right] = tar[right], tar[left]
		}
	}
	// 把哨兵放到最终满足左边都小于他，右边都大于他的点
	tar[low], tar[left] = tar[left], tar[low]
	return left
}