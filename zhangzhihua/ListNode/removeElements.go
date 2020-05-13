package ListNode

//203. 移除链表元素
//删除链表中等于给定值 val 的所有节点。
//
//示例:
//
//输入: 1->2->6->3->4->5->6, val = 6
//输出: 1->2->3->4->5

//执行用时 :
//8 ms
//, 在所有 Go 提交中击败了
//91.33%
//的用户
//内存消耗 :
//4.7 MB
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
func removeElements(head *ListNode, val int) *ListNode {

	node := new(ListNode)
	node.Next = head
	pre, cur := node, head
	for cur != nil {
		if cur.Val != val {
			pre = cur
			cur = cur.Next
		} else {
			pre.Next = cur.Next
			cur = cur.Next
		}
	}
	return node.Next
}
