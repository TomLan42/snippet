/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	inStack  []int
	outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for len(this.outStack) > 0 {
		val := this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
		this.inStack = append(this.inStack, val)
	}
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	for len(this.inStack) > 0 {
		val := this.inStack[len(this.inStack)-1]
		this.inStack = this.inStack[:len(this.inStack)-1]
		this.outStack = append(this.outStack, val)
	}

	if len(this.outStack) == 0 {
		return 0
	}

	val := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return val

}

/** Get the front element. */
func (this *MyQueue) Peek() int {

	for len(this.inStack) > 0 {
		val := this.inStack[len(this.inStack)-1]
		this.inStack = this.inStack[:len(this.inStack)-1]
		this.outStack = append(this.outStack, val)
	}

	if len(this.outStack) == 0 {
		return 0
	}

	val := this.outStack[len(this.outStack)-1]
	return val

}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.outStack) == 0 && len(this.inStack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

