package stack

import (
	"testing"
)

func TestRemoveOuterParentheses(t *testing.T) {
	//S := "(()())(())"
	S := "(()())(())(()(()))"

	t.Log(removeOuterParentheses(S))
}
