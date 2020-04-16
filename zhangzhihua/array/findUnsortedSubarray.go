package array

//执行用时 :
//1252 ms
//, 在所有 Go 提交中击败了
//5.26%
//的用户
//内存消耗 :
//6.1 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//581. 最短无序连续子数组
//给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
//你找到的子数组应是最短的，请输出它的长度。
//
//示例 1:
//
//输入: [2, 6, 4, 8, 10, 9, 15]
//输出: 5
//解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//说明 :
//
//输入的数组长度范围在 [1, 10,000]。
//输入的数组可能包含重复元素 ，所以升序的意思是<=。

func findUnsortedSubarray(nums []int) int {

	o := order(nums)

	l := -1
	r := -1

	for i := 0; i < len(nums); i++ {
		if o[i] != nums[i] {
			l = i
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if o[i] != nums[i] {
			r = i
			break
		}
	}

	if l == -1 || r == -1 {
		return 0
	}

	return r - l + 1

}

func order(ary []int) []int {
	newAry := make([]int, len(ary)) //使用copy函数必须复制切片的结构必须和源数据结构一致
	copy(newAry, ary)
	for i := 0; i < len(newAry); i++ {
		for j := i + 1; j <= len(newAry)-1; j++ {
			if newAry[i] > newAry[j] {
				newAry[i], newAry[j] = newAry[j], newAry[i]
			}
		}
	}
	return newAry
}
func min(one, two int) int {
	if one > two {
		return two
	}
	return one
}
