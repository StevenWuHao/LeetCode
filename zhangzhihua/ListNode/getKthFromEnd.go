package ListNode

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2.3 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//面试题22. 链表中倒数第k个节点
//输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。
//
//
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 k = 2.
//
//返回链表 4->5.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var listNodeCnt int
	var listNode []*ListNode
	for head != nil {
		listNodeCnt++
		listNode = append(listNode, head)
		head = head.Next
	}
	return listNode[listNodeCnt-k]
}
