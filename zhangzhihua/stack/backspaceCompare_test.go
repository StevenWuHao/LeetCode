package stack

import "testing"

func TestBackspaceCompare(t *testing.T) {
	S := "ab#c"
	T := "ad#c"

	t.Log(backspaceCompare(S, T))
}
