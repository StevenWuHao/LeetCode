package tree

//执行用时 :
//4 ms
//, 在所有 Go 提交中击败了
//94.62%
//的用户
//内存消耗 :
//4.4 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
//
//例如：
//
//给定二叉树 [3,9,20,null,null,15,7]，
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最大深度 3 。
//
// 
//
//提示：
//
//节点总数 <= 10000
//注意：本题与主站 104 题相同：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxDepth(root *TreeNode) int {

	depth := 0
	if root == nil {
		return 0
	}
	var queue []*TreeNode

	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			left := queue[i].Left
			right := queue[i].Right
			if left != nil {
				queue = append(queue, left)
			}

			if right != nil {
				queue = append(queue, right)
			}
		}
		queue = queue[qLen:]
		depth++
	}

	return depth
}
