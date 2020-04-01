package tree

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//5.1 MB
//, 在所有 Go 提交中击败了
//87.32%
//的用户

//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最小深度  2.
//
//通过次数62,655提交次数149,662
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func minDepth(root *TreeNode) int {
	min := 0

	if root == nil {
		return min
	}
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		min++
		for i := 0; i < qLen; i++ {
			if queue[i].Right == nil && queue[i].Left == nil {
				return min
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[qLen:]
	}

	return min
}
