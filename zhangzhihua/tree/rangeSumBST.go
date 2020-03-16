package tree

//执行用时 :
//100 ms
//, 在所有 Go 提交中击败了
//88.06%
//的用户
//内存消耗 :
//7.9 MB
//, 在所有 Go 提交中击败了
//45.45%
//的用户

//给定二叉搜索树的根结点 root，返回 L 和 R（含）之间的所有结点的值的和。
//
//二叉搜索树保证具有唯一的值。
//
// 
//
//示例 1：
//
//输入：root = [10,5,15,3,7,null,18], L = 7, R = 15
//输出：32
//示例 2：
//
//输入：root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
//输出：23
// 
//
//提示：
//
//树中的结点数量最多为 10000 个。
//最终的答案保证小于 2^31。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/range-sum-of-bst
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func rangeSumBST(root *TreeNode, L int, R int) int {

	sum := 0
	if root == nil || (L == 0 && R == 0) {
		return sum
	}

	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	sum += rangeSumBST(root.Left, L, R)
	sum += rangeSumBST(root.Right, L, R)

	return sum
}
