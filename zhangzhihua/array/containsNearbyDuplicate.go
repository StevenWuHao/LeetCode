package array

//执行用时 :
//824 ms
//, 在所有 Go 提交中击败了
//16.23%
//的用户
//内存消耗 :
//5.1 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//219. 存在重复元素 II
//给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。
//
//
//
//示例 1:
//
//输入: nums = [1,2,3,1], k = 3
//输出: true
//示例 2:
//
//输入: nums = [1,0,1,1], k = 1
//输出: true
//示例 3:
//
//输入: nums = [1,2,3,1,2,3], k = 2
//输出: false
//通过次数40,924提交次数107,145
//在真实的面试中遇到过这道题？

func containsNearbyDuplicate(nums []int, k int) bool {

	nLen := len(nums)

	for i := 0; i < nLen; i++ {
		min := min(nLen, i+k+1)
		for j := i + 1; j < min; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func min(one, two int) int {
	if one > two {
		return two
	}
	return one
}
