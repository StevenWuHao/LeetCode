package tree

import "fmt"

//执行用时 :
//12 ms
//, 在所有 Go 提交中击败了
//48.28%
//的用户
//内存消耗 :
//5.7 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。
//
//
//示例 1:
//给定二叉树 [3,9,20,null,null,15,7]
//3
/// \
//9  20
///  \
//15   7
//返回 true 。
//示例 2:
//给定二叉树 [1,2,2,3,3,null,null,4,4]
//1
/// \
//2   2
/// \
//3   3
/// \
//4   4
//返回 false 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/check-balance-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	l := deep(root.Left, 0)
	r := deep(root.Right, 0)
	tmp := 0

	fmt.Println(l)
	fmt.Println(r)
	if l > r {
		tmp = l - r
	} else {
		tmp = r - l
	}

	if tmp > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func deep(root *TreeNode, deth int) int {
	if root == nil {
		return deth
	}
	l := deep(root.Left, deth+1)
	r := deep(root.Right, deth+1)
	return max(l, r)
}

func max(l int, r int) int {
	if l > r {
		return l
	}
	return r
}
