package main

import (
	"fmt"
)

func main() {
	fmt.Println("test here!")
	isIsomorphic("agg", "edd")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	inner := func(s string, t string) bool {
		s2t := make(map[uint8]uint8)
		for i := 0; i < len(s); i++ {
			if after, ok := s2t[s[i]]; ok {
				if after != t[i] {
					return false
				}
			} else {
				s2t[s[i]] = t[i]
			}
		}
		return true
	}
	return inner(s, t) && inner(t, s)
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
