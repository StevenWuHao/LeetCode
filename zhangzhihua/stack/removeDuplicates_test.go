package stack

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	//static := []int{1, 2, 3}
	//fmt.Println(static[:len(static)-1])

	S := "abbaca"
	r := removeDuplicates(S)
	fmt.Println(r)
}
