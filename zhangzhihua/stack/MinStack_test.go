package stack

import (
	"fmt"
	"testing"
)

func TestIndex(T *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	obj.Pop()
	param3 := obj.Top()
	param4 := obj.GetMin()

	fmt.Println(param3)
	fmt.Println(param4)
}
