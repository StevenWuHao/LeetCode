package stack


//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2 MB
//, 在所有 Go 提交中击败了
//60.00%
//的用户

//使用栈实现队列的下列操作：
//
//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。
//示例:
//
//MyQueue queue = new MyQueue();
//
//queue.push(1);
//queue.push(2);
//queue.peek();  // 返回 1
//queue.pop();   // 返回 1
//queue.empty(); // 返回 false
//说明:
//
//你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
//你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MyQueue struct {
	in  []int
	out []int
}

/** Initialize your data structure here. */
func NewConstructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.flip()
	if this.Empty() {
		return 0
	}
	top := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return top
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.flip()
	top := 0
	if len(this.out) > 0 {
		top = this.out[len(this.out)-1]
	}
	return top
}

func (this *MyQueue) flip() {
	if len(this.out) != 0 {
		return
	}
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i]) //将入栈的数据，反过来放到出栈
		this.in = this.in[:i]
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.in) > 0 || len(this.out) > 0 {
		return false
	}
	return true
}