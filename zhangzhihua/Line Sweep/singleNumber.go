package Line_Sweep

import "fmt"

//排序法
func singleNumber(nums []int) int {

	//先排序
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	//排序后，如果是1对的一定是连续的，否则直接返回
	cur := 0
	for cur < len(nums)-2 {
		if nums[cur] != nums[cur+1] {
			return nums[cur]
		}
		cur += 2
	}
	return nums[cur]
}

//map法
func singleNumberx(nums []int) int {

	var m = make(map[int]int)
	for _, val := range nums {
		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}
	fmt.Println(m)

	for key, val := range m {
		if val == 1 {
			return key
		}
	}

	return 0
}
