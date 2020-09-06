package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("test here!")
	//m := []int{1,2,3}
	//n := make([][]int, 0)
	//n = append(n, m)
	//n = append(n, m)
	//n = append(n, m)
	//fmt.Println(n)
	//
	//fmt.Println(kthSmallest(n, 3))
	t := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(binaryTreePaths(&t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	tmp := make([]string, 0)
	helper(root, tmp)
	return res
}

func helper(root *TreeNode, now []string) {
	if root.Left != nil {
		helper(root.Left, append(now, strconv.Itoa(root.Val)))
	}
	if root.Right != nil {
		helper(root.Right, append(now, strconv.Itoa(root.Val)))
	}
	now = append(now, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		res = append(res, strings.Join(now, "->"))
		return
	}
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
