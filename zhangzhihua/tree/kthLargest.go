package tree

//执行用时 :
//12 ms
//, 在所有 Go 提交中击败了
//85.71%
//的用户
//内存消耗 :
//6.4 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//给定一棵二叉搜索树，请找出其中第k大的节点。
//
// 
//
//示例 1:
//
//输入: root = [3,1,4,null,2], k = 1
//   3
//  / \
// 1   4
//  \
//   2
//输出: 4
//示例 2:
//
//输入: root = [5,3,6,2,4,null,null,1], k = 3
//       5
//      / \
//     3   6
//    / \
//   2   4
//  /
// 1
//输出: 4
// 
//
//限制：
//
//1 ≤ k ≤ 二叉搜索树元素个数
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var stack []*TreeNode
	count := 0
	curNode := root
	var res []int

	for curNode != nil || len(stack) > 0 {
		if curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
			continue
		}

		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		curNode = node.Right
		count++
	}
	return res[len(res)-k-1]
}
