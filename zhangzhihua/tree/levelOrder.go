package tree

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//3 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
//
// 
//
//例如:
//给定二叉树: [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其层次遍历结果：
//
//[
//[3],
//[9,20],
//[15,7]
//]
// 
//
//提示：
//
//节点总数 <= 1000
//注意：本题与主站 102 题相同：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	var queue [] TreeNode

	if root == nil {
		return ret
	}

	queue = append(queue, *root)

	for len(queue) > 0 {
		var tmp []int
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			node := queue[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, *node.Left)
			}

			if node.Right != nil {
				queue = append(queue, *node.Right)
			}
		}
		queue = queue[queueLen:]
		ret = append(ret, tmp)
	}

	return ret
}
