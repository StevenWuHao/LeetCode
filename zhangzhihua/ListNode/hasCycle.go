package ListNode

//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//
//
//示例 1：
//
//输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
//示例 2：
//
//输入：head = [1,2], pos = 0
//输出：true
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
//示例 3：
//
//输入：head = [1], pos = -1
//输出：false
//解释：链表中没有环。
//
//
//
//
//进阶：
//
//你能用 O(1)（即，常量）内存解决此问题吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/linked-list-cycle
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//执行用时 :
//12 ms
//, 在所有 Go 提交中击败了
//21.72%
//的用户
//内存消耗 :
//6.1 MB
//, 在所有 Go 提交中击败了
//11.76%
//的用

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	var m = make(map[*ListNode]int)

	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = head.Val
		head = head.Next
	}
	return false
}

/**
设置了一个永远不会遇到的值，如果遇到说明是一个圆
*/

//执行用时 :
//8 ms
//, 在所有 Go 提交中击败了
//81.53%
//的用户
//内存消耗 :
//3.8 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户
func hasCyclex(head *ListNode) bool {

	unKnown := 123123123131

	for head != nil {
		if head.Val == unKnown {
			return true
		}
		head.Val = unKnown
		head = head.Next
	}

	return false
}
