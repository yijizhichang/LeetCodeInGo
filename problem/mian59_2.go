package problem

/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

*/
type MaxQueue struct {
	q1  []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q1 = append(this.q1, value)
	for len(this.max) != 0 && value > this.max[len(this.max)-1] {
		// 如果入队的从在max队列中从小到大地比较，把比他小的都踢出队列，再把当前的值放入max队列中
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.q1) != 0 {
		n = this.q1[0]
		this.q1 = this.q1[1:]
		// 如果出队的是最大的值，则把最大的在最大队列中扔出去
		if this.max[0] == n {
			this.max = this.max[1:]
		}
	}
	return n
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
