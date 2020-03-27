package tree

import (
	"math"
)

//执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户 内存消耗 : 2.6 MB ,

//在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
//
//如果二叉树的两个节点深度相同，但父节点不同，则它们是一对堂兄弟节点。
//
//我们给出了具有唯一值的二叉树的根节点 root，以及树中两个不同节点的值 x 和 y。
//
//只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true。否则，返回 false。
//
//
//
//示例 1：
//
//
//输入：root = [1,2,3,4], x = 4, y = 3
//输出：false
//示例 2：
//
//
//输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
//输出：true
//示例 3：
//
//
//
//输入：root = [1,2,3,null,4], x = 2, y = 3
//输出：false
//
//
//提示：
//
//二叉树的节点数介于 2 到 100 之间。
//每个节点的值都是唯一的、范围为 1 到 100 的整数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/cousins-in-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func isCousins(root *TreeNode, x int, y int) bool {

	if root == nil || x == y {
		return false
	}

	var queue []*TreeNode
	queue = append(queue, root)
	deep := 1

	for len(queue) > 0 {
		qLen := len(queue)
		xIdx := -1
		yIdx := -1
		for i := 0; i < qLen; i++ {
			node := queue[i]
			if node != nil {
				if node.Val == x {
					xIdx = i
				}
				if node.Val == y {
					yIdx = i
				}
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}

		if xIdx != -1 && yIdx != -1 {
			cha := math.Abs(float64(xIdx - yIdx))
			if cha > 1 {
				return true
			}

			//小的索引是奇数，返回true
			if cha == 1 {
				min := 0
				if xIdx > yIdx {
					min = yIdx
				} else {
					min = xIdx
				}

				if min%2 != 0 {
					return true
				}
			}
			return false
		}
		deep++
		queue = queue[qLen:]
	}

	return false
}
