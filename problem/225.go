package problem

// 用队列实现栈

type MyStack struct {
	slice []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		slice: make([]int, 0),
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.slice = append(this.slice, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ret := 0
	if !this.Empty() {
		ret = this.slice[len(this.slice)-1]
		this.slice = this.slice[:len(this.slice)-1]
	}
	return ret
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	} else {
		return this.slice[len(this.slice)-1]
	}
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.slice) <= 0 {
		return true
	} else {
		return false
	}
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */