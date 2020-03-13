package tree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tn := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9, Left: nil, Right: nil},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}}

	//fmt.Println(tn)
	fmt.Println(levelOrderBottom(tn))
}
