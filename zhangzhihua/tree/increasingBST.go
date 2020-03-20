package tree

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2.2 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var stack []*TreeNode

	var res []int
	node := root
	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
			continue
		}

		n := stack[len(stack)-1]
		res = append(res, n.Val)
		stack = stack[:len(stack)-1]
		node = n.Right
	}

	pre := &TreeNode{res[0], nil, nil}
	tmp := pre
	for _, v := range res[1:] {
		next := &TreeNode{v, nil, nil}
		tmp.Right = next
		tmp = next
	}
	return pre
}
