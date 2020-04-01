package tree

//执行用时 :
//4 ms
//, 在所有 Go 提交中击败了
//72.43%
//的用户
//内存消耗 :
//3 MB
//, 在所有 Go 提交中击败了
//6.00%
//的用户

//101. 对称二叉树
//给定一个二叉树，检查它是否是镜像对称的。
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//说明:
//
//如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
func c(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var lQueues []*TreeNode
	var rQueues []*TreeNode

	lQueues = append(lQueues, root)
	rQueues = append(rQueues, root)

	for len(lQueues) > 0 {
		qLen := len(lQueues)
		for i := 0; i < qLen; i++ {

			if lQueues[i].Val != rQueues[i].Val {
				return false
			}

			//如果左节点不为nil，对称的右节点也不能为nil
			if lQueues[i].Left != nil {
				if rQueues[i].Right == nil {
					return false
				}
				lQueues = append(lQueues, lQueues[i].Left)
			}

			if lQueues[i].Right != nil {
				if rQueues[i].Left == nil {
					return false
				}
				lQueues = append(lQueues, lQueues[i].Right)
			}

			if rQueues[i].Right != nil {
				rQueues = append(rQueues, rQueues[i].Right)
			}
			if rQueues[i].Left != nil {
				rQueues = append(rQueues, rQueues[i].Left)
			}
		}
		lQueues = lQueues[qLen:]
		rQueues = rQueues[qLen:]
	}
	return true
}

/**
我觉得代码方法没问题，但是超过时间限制
*/
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var queue []*TreeNode
	NotNullTreeNodeCnt := 0
	queue = append(queue, root)
	NotNullTreeNodeCnt++
	for len(queue) > 0 && NotNullTreeNodeCnt > 0 {
		qLen := len(queue)
		NotNullTreeNodeCnt = 0
		lp := 0
		rp := qLen - 1
		for lp < rp {
			if queue[lp].Val != queue[rp].Val {
				return false
			}
			lp++
			rp--
		}
		for i := 0; i < qLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				NotNullTreeNodeCnt++
			} else {
				queue = append(queue, &TreeNode{-1, nil, nil})
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				NotNullTreeNodeCnt++
			} else {
				queue = append(queue, &TreeNode{-1, nil, nil})
			}
		}
		if NotNullTreeNodeCnt == 0 {
			break
		}

		queue = queue[qLen:]
	}
	return true
}
