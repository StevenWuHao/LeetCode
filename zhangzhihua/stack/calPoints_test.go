package stack

import (
	"fmt"
	"testing"
)

func TestCalPoints(t *testing.T) {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	r := calPoints(ops)

	fmt.Println(r)
}
