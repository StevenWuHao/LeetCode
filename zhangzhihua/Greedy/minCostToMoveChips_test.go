package Greedy

import (
	"fmt"
	"testing"
)

func TestMinCostToMoveChips(t *testing.T) {

	chips := []int{1, 2, 2, 2, 2}
	cnt := minCostToMoveChips(chips)

	fmt.Println(cnt)
}
