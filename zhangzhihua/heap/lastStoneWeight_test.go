package heap

import (
	"fmt"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {

	stones := []int{2, 7, 4, 1, 8, 1}
	lastNumber := lastStoneWeight(stones)
	fmt.Println(lastNumber)
}
