package tree

//执行用时 :
//4 ms
//, 在所有 Go 提交中击败了
//75.00%
//的用户
//内存消耗 :
//3.9 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//给定一个 N 叉树，返回其节点值的后序遍历。
//
//例如，给定一个 3叉树 :
//
//
//
//
//
//
//
//返回其后序遍历: [5,6,3,2,4,1].
//
//
//
//说明: 递归法很简单，你可以使用迭代法完成此题吗?
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {

	if root == nil {
		return nil
	}
	var stack []*Node
	var res []int

	//把头结点放到列队和返回数据中
	stack = append(stack, root)

	for len(stack) > 0 {
		QLen := len(stack)
		node := stack[len(stack)-1] //先进先出
		stack = stack[:QLen-1]      //先进先出
		if node.Children != nil {
			for k := 0; k < len(node.Children); k++ {
				stack = append(stack, node.Children[k])
			}
		}
		res = append(res, node.Val)
	}

	if res == nil {
		return res
	}

	lp := 0
	rp := len(res) - 1
	for lp < rp { //指针相撞
		res[lp], res[rp] = res[rp], res[lp]
		lp++
		rp--
	}

	return res
}

//func echo(q []*Node) {
//	for _, v := range q {
//		fmt.Printf("当前q的长度 %d,值为%d \n\r", len(q), v.Val)
//	}
//}
//
//func echoChildren(node *Node) {
//	for _, v := range node.Children {
//		fmt.Printf("值为%d \n\r", v.Val)
//	}
//	fmt.Printf("\n\r")
//}
