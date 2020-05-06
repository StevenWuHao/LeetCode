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

//面试题 02.02. 返回倒数第 k 个节点
//实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
//
//注意：本题相对原题稍作改动
//
//示例：
//
//输入： 1->2->3->4->5 和 k = 2
//输出： 4
//说明：
//
//给定的 k 保证是有效的。
//
//通过次数9,603提交次数12,175

func kthToLast(head *ListNode, k int) int {
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	return vals[len(vals)-k]
}
