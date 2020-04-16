package array

import (
	"math"
)

//执行用时 :
//4 ms
//, 在所有 Go 提交中击败了
//92.12%
//的用户
//内存消耗 :
//3 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
//
//示例 1:
//
//输入: [3, 2, 1]
//
//输出: 1
//
//解释: 第三大的数是 1.
//示例 2:
//
//输入: [1, 2]
//
//输出: 2
//
//解释: 第三大的数不存在, 所以返回最大的数 2 .
//示例 3:
//
//输入: [2, 2, 3, 1]
//
//输出: 1
//
//解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
//存在两个值为2的数，它们都排第二。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/third-maximum-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func thirdMax(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[1]
	}

	max := nums[0]
	mid := math.MinInt32 - 1
	min := math.MinInt32 - 1

	for i := 1; i < length; i++ {

		if nums[i] == max || nums[i] == mid || nums[i] == min {
			continue
		}

		if nums[i] > max {
			min = mid
			mid = max
			max = nums[i]
			continue
		}
		if nums[i] > mid {
			min = mid
			mid = nums[i]
			continue
		}
		if nums[i] > min {
			min = nums[i]
			continue
		}
	}
	if min != math.MinInt32-1 {
		return min
	}

	return max
}
