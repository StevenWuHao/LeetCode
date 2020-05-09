package ListNode

import (
	"fmt"
	"testing"
)

//理解链表内存
func TestReverseList(t *testing.T) {

	var head = &ListNode{0, nil}
	first := &ListNode{1, nil}
	two := &ListNode{2, nil}
	three := &ListNode{3, nil}
	head.Next = first
	first.Next = two
	two.Next = three

	a := head
	b := head.Next.Next
	list := reverseList(a)
	node := list
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

	for b != nil {
		fmt.Println(b.Val)
		b = b.Next
	}
}
