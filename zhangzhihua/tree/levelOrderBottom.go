package tree

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//3 MB
//, 在所有 Go 提交中击败了
//70.59%
//的用户

//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其自底向上的层次遍历为：
//
//[
//[15,7],
//[9,20],
//[3]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func levelOrderBottom(root *TreeNode) [][]int {

	var ret [][]int
	var queue [] *TreeNode

	if root == nil {
		return ret
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		var tmp []int
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			tmp = append(tmp, queue[i].Val)

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		ret = append(ret, tmp)
		queue = queue[queueLen:]
	}

	lp := 0
	rp := len(ret) - 1
	for lp < rp {
		ret[lp], ret[rp] = ret[rp], ret[lp]
		lp++
		rp--
	}

	return ret
}
