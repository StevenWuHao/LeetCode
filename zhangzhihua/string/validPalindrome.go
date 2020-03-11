package string

import "fmt"

//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//
//示例 1:
//
//输入: "aba"
//输出: True
//示例 2:
//
//输入: "abca"
//输出: True
//解释: 你可以删除c字符。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-palindrome-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

	//s := "aba"
	s := "abca"
	//s := "abcba"
	//s := "abc"
	//s := "aba"
	//s := "cxcaac"

	//执行用时 :
	//	16 ms
	//	, 在所有 golang 提交中击败了
	//	95.06%
	//		的用户
	//内存消耗 :
	//	6.3 MB
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户
	fmt.Println(validPalindrome(s))
}
func validPalindrome(s string) bool {

	lp := 0
	rp := len(s) - 1

	//先检查整个是不是回文
	if isHuiWen(s, lp, rp) {
		return true
	}

	//尝试删除
	for lp < rp {

		if s[lp] == s[rp] {
			lp++
			rp--
			continue
		}
		return isHuiWen(s, lp+1, rp) || isHuiWen(s, lp, rp-1)

	}
	return false
}

func isHuiWen(s string, i int, j int) bool {

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
