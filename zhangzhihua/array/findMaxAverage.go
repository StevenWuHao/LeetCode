package array

import "fmt"

//执行用时 :
//764 ms
//, 在所有 Go 提交中击败了
//21.13%
//的用户
//内存消耗 :
//6.6 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。
//
//示例 1:
//
//输入: [1,12,-5,-6,50,3], k = 4
//输出: 12.75
//解释: 最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
//
//
//注意:
//
//1 <= k <= n <= 30,000。
//所给数据范围 [-10,000，10,000]。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-average-subarray-i
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findMaxAverage(nums []int, k int) float64 {

	var max int
	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {
		var sum int
		maxLen := i + k
		if maxLen > numsLen {
			break
		}
		for j := i; j < maxLen; j++ {
			sum += nums[j]
			fmt.Printf("num[j]= %d ", nums[j])
		}
		fmt.Printf("sum = %d \n\r", sum)
		if sum > max && max == 0 {
			max = sum
		}
	}

	return float64(max) / float64(k)
}
