package problem

import "container/heap"

/*
295.数据流的中位数
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。

示例：

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
*/

type IntPQ []int

func (pq IntPQ) Len() int { return len(pq) }

func (pq IntPQ) Less(i int, j int) bool { return pq[i] < pq[j] }

func (pq IntPQ) Swap(i int, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *IntPQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *IntPQ) Pop() interface{} {
	old := *pq

	x := old[len(old)-1]
	*pq = old[:len(old)-1]

	return x
}

type MedianFinder struct {
	lows  IntPQ
	highs IntPQ
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.lows.Len() == this.highs.Len() {
		heap.Push(&this.highs, num)
	} else {
		heap.Push(&this.lows, -num)
	}

	// re-balance
	if this.lows.Len() == 0 {
		return
	}

	x, y := -this.lows[0], this.highs[0]
	if x > y {
		heap.Pop(&this.lows)
		heap.Pop(&this.highs)

		heap.Push(&this.lows, -y)
		heap.Push(&this.highs, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.highs.Len() > this.lows.Len() {
		return float64(this.highs[0])
	}

	return (float64(this.highs[0]) - float64(this.lows[0])) / 2
}
