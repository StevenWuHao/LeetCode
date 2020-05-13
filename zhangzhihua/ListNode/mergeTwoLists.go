package ListNode

//21. 合并两个有序链表
//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

//执行用时 :
//4 ms
//, 在所有 Go 提交中击败了
//66.62%
//的用户
//内存消耗 :
//2.5 MB
//, 在所有 Go 提交中击败了
//72.73%
//的用户

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	node := new(ListNode)
	pre := node

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l2 == nil {
		pre.Next = l1
	}
	if l1 == nil {
		pre.Next = l2
	}

	return node.Next
}
