/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

package code

// @lc code=start
type MyStack struct {
	Items []int
}

func Constructor() MyStack {
	return MyStack{Items: []int{}}
}

func (this *MyStack) Push(x int) {
	this.Items = append(this.Items, x)
	return
}

func (this *MyStack) Pop() int {
	last := this.Items[len(this.Items)-1]
	this.Items = this.Items[:len(this.Items)-1]
	return last
}

func (this *MyStack) Top() int {
	return this.Items[len(this.Items)-1]
}

func (this *MyStack) Empty() bool {
	if this.Items == nil || len(this.Items) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
