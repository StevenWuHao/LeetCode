package sort

//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [9,4]
//说明:
//
//输出结果中的每个元素一定是唯一的。
//我们可以不考虑输出结果的顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//3.2 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户
func intersection(nums1 []int, nums2 []int) []int {

	nums1Len := len(nums1)
	nums2Len := len(nums2)

	var processNums []int
	var anthor []int
	var result []int
	dictionary := make(map[int]int)
	if nums1Len > nums2Len {
		processNums = nums1
		anthor = nums2
	} else {
		processNums = nums2
		anthor = nums1
	}

	for _, v := range processNums {
		if _, ok := dictionary[v]; !ok {
			dictionary[v] = 1
		}
	}

	for _, v := range anthor {
		if _, ok := dictionary[v]; ok {
			result = append(result, v)
			delete(dictionary, v)
		}
	}

	return result
}
