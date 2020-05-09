package ListNode

//执行用时 :
//8 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//6.1 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//面试题 02.01. 移除重复节点
//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
//
//示例1:
//
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
//示例2:
//
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
//提示：
//
//链表长度在[0, 20000]范围内。
//链表元素在[0, 20000]范围内。
//进阶：
//
//如果不得使用临时缓冲区，该怎么解决？

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	var mp = make(map[int]int)

	cur := head
	var pre *ListNode = nil
	for cur != nil {
		_, ok := mp[cur.Val]
		if !ok {
			mp[cur.Val] = cur.Val
			pre = cur
			cur = cur.Next
			continue
		}
		pre.Next = cur.Next
		cur = pre.Next
	}
	return head
}
