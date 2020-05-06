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
