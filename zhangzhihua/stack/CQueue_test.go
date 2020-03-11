package stack

import "testing"

func TestCQueue(t *testing.T) {

	q := ConstructorA()
	q.AppendTail(4)
	q.AppendTail(3)
	t.Log(q.DeleteHead())
	t.Log(q.DeleteHead())
	t.Log(q.DeleteHead())
}
