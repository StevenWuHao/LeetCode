package tree

import (
	"fmt"
	"testing"
)

func TestIsCousins(t *testing.T) {

	four := &TreeNode{4, nil, nil}

	six := &TreeNode{6, nil, nil}
	three := &TreeNode{3, four, nil}

	two := &TreeNode{2, three, nil}
	five := &TreeNode{5, nil, six}

	one := &TreeNode{1, two, five}

	bl := isCousins(one, 3, 6)
	fmt.Println(bl)
}
