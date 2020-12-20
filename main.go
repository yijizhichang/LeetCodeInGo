package main

import (
	"fmt"
)

func main() {
	fmt.Println("test here!")
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeDuplicateLetters(s string) string {
	shengyu := make([]int32, 26)
	stack := make([]int32, 0)
	// 是否在栈中，存在为true
	var exist [26]bool
	for _, this := range s {
		shengyu[this-'a']++
	}
	for _, this := range s {
		// 如果栈中已有，则不用加
		if exist[this-'a'] {
			shengyu[this-'a']--
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > this && shengyu[stack[len(stack)-1]-'a'] > 0 {
			// 标记为栈中不含栈顶元素
			exist[stack[len(stack)-1]-'a'] = false
			// 删除栈顶元素
			stack = stack[:len(stack)-1]
		}
		// 添加新字符
		stack = append(stack, this)
		// 减少该字符出现次数
		shengyu[this-'a']--
		// 标记栈中有该字符
		exist[this-'a'] = true
	}
	ret := ""
	for _, this := range stack {
		ret += string(this)
	}
	return ret
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0
	num := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			num += i + 1
			j++
		} else {
			i--
		}
	}
	return num >= k
}
