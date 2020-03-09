package stack


//执行用时 :
//40 ms
//, 在所有 Go 提交中击败了
//30.77%
//的用户
//内存消耗 :
//6.4 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
//
//示例:
//
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7
// 
//
//提示：
//
//你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
//
//注意：本题与主站 239 题相同：https://leetcode-cn.com/problems/sliding-window-maximum/
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func maxSlidingWindow(nums []int, k int) []int {

	var maxs []int
	wl := 0
	cnt := len(nums)
	if cnt == 0 {
		return maxs
	}
	for k <= cnt {
		maxs = append(maxs, max(nums[wl:k]))
		wl++
		k++
	}

	return maxs
}

func max(nums []int) int {

	len := len(nums)
	if len == 1 {
		return nums[len-1]
	}

	max := nums[0]
	for len > 0 {
		if nums[len-1] > max {
			max = nums[len-1]
		}
		len--
	}

	return max
}
