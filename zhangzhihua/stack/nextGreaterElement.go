package stack

//执行用时 :
//4 ms
//, 在所有 Go 提交中击败了
//91.55%
//的用户
//内存消耗 :
//3.4 MB
//, 在所有 Go 提交中击败了
//23.53%
//的用户

//给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。
//
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。
//
//示例 1:
//
//输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出: [-1,3,-1]
//解释:
//对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
//对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
//对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
//示例 2:
//
//输入: nums1 = [2,4], nums2 = [1,2,3,4].
//输出: [3,-1]
//解释:
//    对于num1中的数字2，第二个数组中的下一个较大数字是3。
//对于num1中的数字4，第二个数组中没有下一个更大的数字，因此输出 -1。
//注意:
//
//nums1和nums2中所有元素是唯一的。
//nums1和nums2 的数组大小都不超过1000。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/next-greater-element-i
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	var mp = make(map[int]int)
	for i := 0; i < len(nums2); i++ {

		//如果栈为空直接放入栈
		if len(stack) == 0 {
			stack = append(stack, nums2[i])
			continue
		}

		//当前的和栈一个个比对
		for {
			if len(stack) == 0 {
				break
			}
			if nums2[i] > stack[len(stack)-1] {
				mp[stack[len(stack)-1]] = nums2[i]
				stack = stack[:len(stack)-1]
				continue
			} else {
				break
			}
		}
		stack = append(stack, nums2[i])

	}
	for {
		if len(stack) == 0 {
			break
		}
		mp[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	var r []int
	for i := 0; i < len(nums1); i++ {
		m, _ := mp[nums1[i]]
		r = append(r, m)
	}

	return r
}
