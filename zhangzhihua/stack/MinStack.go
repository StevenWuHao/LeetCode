package stack

//执行用时 :
//16 ms
//, 在所有 Go 提交中击败了
//98.31%
//的用户
//内存消耗 :
//7.8 MB
//, 在所有 Go 提交中击败了
//12.96%
//的用户

//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) -- 将元素 x 推入栈中。
//pop() -- 删除栈顶的元素。
//top() -- 获取栈顶元素。
//getMin() -- 检索栈中的最小元素。
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/min-stack
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)

	if len(this.min) > 0 {
		min := this.min[len(this.min)-1]
		if min > x {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, min)
		}
	} else {
		this.min = append(this.min, x)
	}

}

func (this *MinStack) Pop() {
	this.stack = this.stack[0 : len(this.stack)-1]
	this.min = this.min[0 : len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
