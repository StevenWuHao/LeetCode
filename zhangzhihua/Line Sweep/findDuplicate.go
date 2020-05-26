package Line_Sweep

func findDuplicate(nums []int) int {
	hash := make(map[int]bool)
	for i, _ := range nums {
		if _, ok := hash[nums[i]]; ok {
			return nums[i]
		} else {
			hash[nums[i]] = true
		}
	}
	return 0
}
