package ListNode

import (
	"fmt"
	"math"
)

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2.2 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。
//
//请你返回该链表所表示数字的 十进制值 。
//
//
//
//示例 1：
//
//
//
//输入：head = [1,0,1]
//输出：5
//解释：二进制数 (101) 转化为十进制数 (5)
//示例 2：
//
//输入：head = [0]
//输出：0
//示例 3：
//
//输入：head = [1]
//输出：1
//示例 4：
//
//输入：head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
//输出：18880
//示例 5：
//
//输入：head = [0,0]
//输出：0
//
//
//提示：
//
//链表不为空。
//链表的结点总数不超过 30。
//每个结点的值不是 0 就是 1。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {

	var binary []int
	for head != nil {
		binary = append(binary, head.Val)
		head = head.Next
	}
	fmt.Printf("cur node val = %d \n\r", binary)
	sum := 0
	i := 0
	j := len(binary) - 1
	for i < len(binary) && j >= 0 {
		sum += binary[i] * int(math.Pow(2, float64(j)))
		i++
		j--
	}

	return sum
}
