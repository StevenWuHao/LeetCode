package tree

//执行用时 :
//32 ms
//, 在所有 Go 提交中击败了
//47.74%
//的用户
//内存消耗 :
//6.6 MB
//, 在所有 Go 提交中击败了
//84.62%
//的用户

//给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
//
//例如，
//
//给定二叉搜索树:
//
//        4
//       / \
//      2   7
//     / \
//    1   3
//
//和值: 2
//你应该返回如下子树:
//
//      2
//     / \
//    1   3
//在上述示例中，如果要找的值是 5，但因为没有节点值为 5，我们应该返回 NULL。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-in-a-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}

	return root
}
