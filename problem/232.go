package problem

// 用栈实现队列

type MyQueue struct {
	slice []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		slice: make([]int, 0),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.slice = append(this.slice, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ret := 0
	if !this.Empty() {
		ret = this.slice[0]
		this.slice = this.slice[1:]
	}
	return ret
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	ret := 0
	if !this.Empty() {
		ret = this.slice[0]
	}
	return ret
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.slice) <= 0 {
		return true
	} else {
		return false
	}
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
