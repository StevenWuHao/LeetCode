package stack

import "testing"

func TestMax(t *testing.T) {
	nums := []int{10, 5, 7}
	max := max(nums)
	t.Log(max)
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	t.Log(maxSlidingWindow(nums, k))
}
