<?php

/**
 * 225. 用队列实现栈
使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
 *
 * https://leetcode-cn.com/problems/implement-stack-using-queues/
 *
 * Class MyStack
 */
class MyStack {
    private $queue = [];
    private $count = 0;
    /**
     * 记录当前队列的索引位置
     * @var int
     */
    private $top_index = 0;

    /**
     * Initialize your data structure here.
     */
    function __construct() {

    }

    /**
     * Push element x onto stack.
     * @param Integer $x
     * @return NULL
     */
    function push($x) {
        $this->queue[] = $x;
        $this->count++;
        if ($this->count > 1) {
            // 栈是先进后出 所以每次入栈都把最后进队列的值移到首位
            for ($i=$this->count-1;$i>0;$i--) {
                $tmp = $this->queue[$i-1];
                $this->queue[$i-1] = $this->queue[$i];
                $this->queue[$i] = $tmp;
            }
        }
    }

    /**
     * Removes the element on top of the stack and returns that element.
     * @return Integer
     */
    function pop() {
        $pop = $this->queue[0];
        $this->top_index++;
        // 每次栈都把队列的第一位移到最后一位
        if ($this->count > 1) {
            for ($i=0;$i<$this->count-1;$i++) {
                $tmp = $this->queue[$i];
                $this->queue[$i] = $this->queue[$i+1];
                $this->queue[$i+1] = $tmp;
            }
        }
        return $pop;
    }

    /**
     * Get the top element.
     * @return Integer
     */
    function top() {
        return $this->queue[0];
    }

    /**
     * Returns whether the stack is empty.
     * @return Boolean
     */
    function empty() {
        return $this->top_index >= $this->count;
    }
}

/**
 * Your MyStack object will be instantiated and called as such:
 * $obj = MyStack();
 * $obj->push($x);
 * $ret_2 = $obj->pop();
 * $ret_3 = $obj->top();
 * $ret_4 = $obj->empty();
 */