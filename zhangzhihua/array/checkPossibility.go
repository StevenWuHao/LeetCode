package array

//执行用时 :
//28 ms
//, 在所有 Go 提交中击败了
//92.96%
//的用户
//内存消耗 :
//6.3 MB
//, 在所有 Go 提交中击败了
//50.00%
//的用户

//all = append(all, []int{4, 2, 3})
//all = append(all, []int{4, 2, 1})
//all = append(all, []int{-1,4,2,3})
//all = append(all, []int{3, 4, 2, 3})
//all = append(all, []int{3, 2, 2, 3, 1})
//all = append(all, []int{3, 2, 2, 3, 4})

//给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
//
//我们是这样定义一个非递减数列的： 对于数组中所有的 i (1 <= i < n)，总满足 array[i] <= array[i + 1]。
//
//
//
//示例 1:
//
//输入: nums = [4,2,3]
//输出: true
//解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
//示例 2:
//
//输入: nums = [4,2,1]
//输出: false
//解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
//
//
//说明：
//
//1 <= n <= 10 ^ 4
//- 10 ^ 5 <= nums[i] <= 10 ^ 5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/non-decreasing-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func checkPossibility(nums []int) bool {

	numsLen := len(nums)
	if numsLen <= 2 {
		return true
	}

	cnt := 0
	for i := numsLen - 2; i >= 0; i-- {

		if cnt > 1 {
			return false
		}

		if nums[i] <= nums[i+1] {
			continue
		}

		if i+2 < numsLen && nums[i] > nums[i+2] {
			nums[i] = nums[i+1]
		}
		cnt++
	}
	return cnt < 2
}
