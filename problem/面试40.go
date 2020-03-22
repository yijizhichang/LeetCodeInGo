package problem

import "container/heap"

/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。



示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]


限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000
 */

func getLeastNumbers(arr []int, k int) []int {
	if k <= 0 {
		return nil
	}
	if k > len(arr) {
		return arr
	}
	ret := make([]int, k-1)
	ret = arr[:k]
	for i:=k; i < len(arr) ; i++{
		max := 0
		this := arr[i]
		// 每次找到k个里最大的
		for j:=1; j < len(ret); j++ {
			if ret[j] > ret[max] {
				max = j
			}
		}
		// 如果这次target比他小，就替换
		if this < ret[max] {
			ret[max] = this
		}
	}
	return ret
}


func getLeastNumbers(arr []int, k int) []int {
	if k <= 0 {
		return nil
	}

	if k > len(arr) {
		return arr
	}

	h := &MaxHeap{}
	*h = append(*h, arr[:k]...)

	heap.Init(h)

	for i := k; i < len(arr); i++ {
		if arr[i] < (*h)[0] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}

	return *h
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}