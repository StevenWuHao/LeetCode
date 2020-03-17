package tree

import (
	"fmt"
	"testing"
)

func TestIsUnivalTree(t *testing.T) {
	tn := &TreeNode{Val: 0}

	fmt.Println(isUnivalTree(tn))
}
