package stack

import (
	"fmt"
	"testing"
)

func TestMyQueue(t *testing.T) {
	queue := NewConstructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	queue.Pop()
	queue.Pop()
	queue.Pop()
	q := queue.Pop()
	fmt.Println(q)

	//bl := queue.Empty()  // 返回 false
	//
	//fmt.Println(peek)
	//fmt.Println(pop)
	//fmt.Println(bl)
}
