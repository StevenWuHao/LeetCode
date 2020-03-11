package stack

//执行用时 :
//240 ms
//, 在所有 Go 提交中击败了
//74.60%
//的用户
//内存消耗 :
//8.1 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
// 
//
//示例 1：
//
//输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
//示例 2：
//
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
//提示：
//
//1 <= values <= 10000
//最多会对 appendTail、deleteHead 进行 10000 次调用
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type CQueue struct {
	stack []int
	sLen  int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stack = append(this.stack, value)
	this.sLen++
}

func (this *CQueue) DeleteHead() int {
	if this.sLen == 0 {
		return -1
	}
	del := this.stack[0]
	this.stack = this.stack[1:]
	this.sLen--
	return del
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
