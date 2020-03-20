package tree

import (
	"fmt"
	"testing"
)

func TestIncreasingBST(t *testing.T) {

	four := &TreeNode{4, nil, nil}
	seven := &TreeNode{7, nil, nil}
	nine := &TreeNode{9, nil, nil}
	eight := &TreeNode{8, seven, nine}
	six := &TreeNode{6, nil, eight}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, one, nil}
	three := &TreeNode{3, two, four}
	five := &TreeNode{5, three, six}

	fmt.Println(increasingBST(five))
}
