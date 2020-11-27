package linked_list

/**
执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
3.1 MB
, 在所有 Go 提交中击败了
69.74%
的用户

从尾到头打印链表

输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

 

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
 

限制：

0 <= 链表长度 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	pre := head
	var cur *ListNode
	for pre != nil {
		tmp := pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
	}

	var arr []int
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	return arr
}