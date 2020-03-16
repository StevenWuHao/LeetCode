package tree

//执行用时 :
//4 ms
//, 在所有 Go 提交中击败了
//93.62%
//的用户
//内存消耗 :
//4.4 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。
//
//示例:
//给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//          0
//         / \
//       -3   9
//       /   /
//     -10  5
//通过次数1,988提交次数2,495
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-height-tree-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
