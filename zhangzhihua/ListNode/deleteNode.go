package ListNode

//执行用时 :
//4 ms
//, 在所有 Go 提交中击败了
//81.47%
//的用户
//内存消耗 :
//2.9 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//面试题 02.03. 删除中间节点
//实现一种算法，删除单向链表中间的某个节点（除了第一个和最后一个节点，不一定是中间节点），假定你只能访问该节点。
//
//
//
//示例：
//
//输入：单向链表a->b->c->d->e->f中的节点c
//结果：不返回任何数据，但该链表变为a->b->d->e->f

//思路简单啊：将后面的内存地址，复制到当前
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	*node = *node.Next
}

//面试题18. 删除链表的节点
//给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
//返回删除后的链表的头节点。
//
//注意：此题对比原题有改动
//
//示例 1:
//
//输入: head = [4,5,1,9], val = 5
//输出: [4,1,9]
//解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//示例 2:
//
//输入: head = [4,5,1,9], val = 1
//输出: [4,5,9]
//解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
//
//
//说明：
//
//题目保证链表中节点的值互不相同
//若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2.9 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

func deleteNodex(head *ListNode, val int) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		if cur.Val == val {
			if cur.Next == nil {
				pre.Next = nil
			} else {
				*cur = *cur.Next
			}
		}
		pre = cur
		cur = cur.Next
	}
	return head
}
