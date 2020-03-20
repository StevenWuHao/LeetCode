package tree

//执行用时 :
//12 ms
//, 在所有 Go 提交中击败了
//64.29%
//的用户
//内存消耗 :
//6.2 MB
//, 在所有 Go 提交中击败了
//58.33%
//的用户

//给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.
//
//示例 1:
//
//输入:
//    3
//   / \
//  9  20
//    /  \
//   15   7
//输出: [3, 14.5, 11]
//解释:
//第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
//注意：
//
//节点值的范围在32位有符号整数范围内。
func averageOfLevels(root *TreeNode) []float64 {
	var ary []float64
	var queue []*TreeNode

	queue = append(queue, root)
	for len(queue) > 0 {
		sLen := len(queue)
		sum := 0
		for i := 0; i < sLen; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[sLen:]
		ary = append(ary, float64(sum)/float64(sLen))
	}

	return ary
}
