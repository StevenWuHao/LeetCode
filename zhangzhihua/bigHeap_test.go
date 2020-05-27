package zhangzhihua

import (
	"fmt"
	"testing"
)

func TestTest(T *testing.T) {
	nums := []int{1, 3, 4, 5, 2, 6, 9, 7, 8, 0}

	for i := len(nums)/2 - 1; i >= 0; i-- {
		initBigHead(i, len(nums), nums)
	}

	for j := len(nums) - 1; j > 0; j-- {
		swap(0, j, nums)
		initBigHead(0, j, nums)
	}
	fmt.Println(nums)
}
