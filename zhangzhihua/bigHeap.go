package zhangzhihua

/**
大顶推和小顶推
*/

func initBigHead(i int, l int, nums []int) {
	left := 2*i + 1
	right := 2*i + 2
	large := i
	if left < l && nums[left] < nums[i] {
		large = left
	}

	if right < l && nums[right] < nums[large] {
		large = right
	}
	if large != i {
		swap(i, large, nums)
		initBigHead(large, l, nums)
	}
}

func swap(i, j int, nums []int) {
	nums[i], nums[j] = nums[j], nums[i]
}
