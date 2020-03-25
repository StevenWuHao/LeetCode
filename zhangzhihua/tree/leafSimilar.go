package tree

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2.5 MB
//, 在所有 Go 提交中击败了
//57.14%
//的用户

//请考虑一颗二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
//
//
//
//举个例子，如上图所示，给定一颗叶值序列为 (6, 7, 4, 9, 8) 的树。
//
//如果有两颗二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。
//
//如果给定的两个头结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。
//
//
//
//提示：
//
//给定的两颗树可能会有 1 到 200 个结点。
//给定的两颗树上的值介于 0 到 200 之间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/leaf-similar-trees
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var r1 []int
	var r2 []int
	leaf(root1, &r1)
	leaf(root2, &r2)

	if len(r1) != len(r2) {
		return false
	}

	index := 0
	for index < len(r1) {
		if r1[index] != r2[index] {
			return false
		}
		index++
	}

	return true
}

func leaf(root *TreeNode, num *[]int) {

	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*num = append(*num, root.Val)
	}

	leaf(root.Left, num)
	leaf(root.Right, num)
}
