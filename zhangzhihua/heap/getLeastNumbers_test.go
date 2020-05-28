package heap

import (
	"fmt"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	nums := []int{0, 1, 2, 1}
	k := 1
	heapSort(nums)
	fmt.Println(nums)
	fmt.Println(nums[:k])
}
