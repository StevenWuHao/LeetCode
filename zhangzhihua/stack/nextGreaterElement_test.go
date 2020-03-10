package stack

import (
	"fmt"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}

	r := nextGreaterElement(nums1, nums2)
	fmt.Println(r)
}
