package tree

import (
	"fmt"
	"testing"
)

func TestPostorder(t *testing.T) {
	five := &Node{5, nil}
	six := &Node{6, nil}
	three := &Node{3, []*Node{five, six}}
	two := &Node{2, nil}
	four := &Node{4, nil}
	root := &Node{1, []*Node{three, two, four}}

	fmt.Println(postorder(root))

}
