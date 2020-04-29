package array

import (
	"fmt"
	"testing"
)

func TestCheckPossibility(t *testing.T) {

	var all [][]int

	all = append(all, []int{4, 2, 3})
	all = append(all, []int{4, 2, 1})
	all = append(all, []int{-1, 4, 2, 3})
	all = append(all, []int{3, 4, 2, 3})
	all = append(all, []int{3, 2, 2, 3, 1})
	all = append(all, []int{3, 2, 2, 3, 4})

	for _, nums := range all {
		fmt.Printf("nums = %d", nums)
		fmt.Printf(" boolean = %t \r\n", checkPossibility(nums))
	}

}
