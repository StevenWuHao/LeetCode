package tree

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2.3 MB
//, 在所有 Go 提交中击败了
//90.91%
//的用户


//如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
//
//只有给定的树是单值二叉树时，才返回 true；否则返回 false。
//
// 
//
//示例 1：
//
//
//
//输入：[1,1,1,1,1,null,1]
//输出：true
//示例 2：
//
//
//
//输入：[2,2,2,5,2]
//输出：false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/univalued-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isUnivalTree(root *TreeNode) bool {

	v := root.Val

	return comp(root, v)
}

func comp(root *TreeNode, v int) bool {

	if root == nil {
		return true
	}

	if root.Val != v {
		return false
	}

	return comp(root.Left, v) && comp(root.Right, v)
}
