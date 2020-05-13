package ListNode

import "fmt"

//面试题 02.06. 回文链表
//编写一个函数，检查输入的链表是否是回文的。
//
//
//
//示例 1：
//
//输入： 1->2
//输出： false
//示例 2：
//
//输入： 1->2->2->1
//输出： true
//
//
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

//执行用时 :
//12 ms
//, 在所有 Go 提交中击败了
//98.15%
//的用户
//内存消耗 :
//7.1 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow := head
	fast := head

	ary := []int{}

	for fast != nil && fast.Next != nil {
		ary = append(ary, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	aryLen := len(ary)
	if aryLen == 0 {
		return true
	}

	if fast != nil {
		slow = slow.Next
	}

	for i := aryLen - 1; i >= 0 && slow != nil; i-- {
		fmt.Printf("ary[i] = %d slow.Val = %d", ary[i], slow.Val)
		if ary[i] != slow.Val {
			return false
		}
		slow = slow.Next
		continue
	}

	return true
}
