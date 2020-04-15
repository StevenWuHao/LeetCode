package array

import "fmt"

//示例有问题，一直提交不过，看了题解，有人评论和我一样，题目有问题

//给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
//
//我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
//
//如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
//
//示例 1:
//
//输入:
//nums = [1, 7, 3, 6, 5, 6]
//输出: 3
//解释:
//索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
//同时, 3 也是第一个符合要求的中心索引。
//示例 2:
//
//输入:
//nums = [1, 2, 3]
//输出: -1
//解释:
//数组中不存在满足此条件的中心索引。
//说明:
//
//nums 的长度范围为 [0, 10000]。
//任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。
//通过次数34,465提交次数93,718
//在真实的面试中遇到过这道题？
//
//是
//
//否
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-pivot-index
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func pivotIndex(nums []int) int {

	numsLen := len(nums)
	if numsLen == 0 || numsLen < 3 {
		return -1
	}

	if numsLen == 1 {
		return 0
	}

	//因为不包含中心索引，初始索引一定是从第2个开始
	for i := 0; i < numsLen-1; i++ {
		//左边总和
		l := sum(nums[:i])
		round := i + 1
		if round > numsLen-1 {
			round = numsLen - 1
		}
		r := sum(nums[round:])
		fmt.Printf("left = %d right = %d \n\r", l, r)
		if l == r {
			return i
		}
	}
	return -1
}

func sum(values []int) int {
	sum := 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum
}
