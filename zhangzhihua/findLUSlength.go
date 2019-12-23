package main

import "fmt"

//给定两个字符串，你需要从这两个字符串中找出最长的特殊序列。最长特殊序列定义如下：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）。
//
//子序列可以通过删去字符串中的某些字符实现，但不能改变剩余字符的相对顺序。空序列为所有字符串的子序列，任何字符串为其自身的子序列。
//
//输入为两个字符串，输出最长特殊序列的长度。如果不存在，则返回 -1。
//
//示例 :
//
//输入: "aba", "cdc"
//输出: 3
//解析: 最长特殊序列可为 "aba" (或 "cdc")
//说明:
//
//两个字符串长度均小于100。
//字符串中的字符仅含有 'a'~'z'。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-uncommon-subsequence-i
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	a := "abc"
	b := "cbc"

	//执行用时 :
	//0 ms
	//, 在所有 golang 提交中击败了
	//100.00%
	//的用户
	//内存消耗 :
	//1.9 MB
	//, 在所有 golang 提交中击败了
	//70.00%
	//的用户

	fmt.Println(findLUSlength(a, b))
}

func findLUSlength(a string, b string) int {

	aL := len(a)
	bL := len(b)

	if a == b {
		return -1
	}

	if aL >= bL {
		return aL
	}

	if aL < bL {
		return bL
	}

	return -1
}
