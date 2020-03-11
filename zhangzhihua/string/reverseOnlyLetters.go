package string

import "fmt"

//给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
//
//
//
//示例 1：
//
//输入："ab-cd"
//输出："dc-ba"
//示例 2：
//
//输入："a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba"
//示例 3：
//
//输入："Test1ng-Leet=code-Q!"
//输出："Qedo1ct-eeLg=ntse-T!"
//
//
//提示：
//
//S.length <= 100
//33 <= S[i].ASCIIcode <= 122
//S 中不包含 \ or "
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-only-letters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

	s := "a-bC-dEf-ghIj"
	//执行用时 :
	//	0 ms
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户
	//内存消耗 :
	//	2 MB
	//	, 在所有 golang 提交中击败了
	//	73.33%
	//		的用户
	fmt.Println(reverseOnlyLetters(s))
}

func reverseOnlyLetters(S string) string {

	lp := 0
	rp := len(S) - 1

	b := []byte(S)
	for lp < rp {

		if !isAZ(S[lp]) {
			lp++
			continue
		}

		if !isAZ(S[rp]) {
			rp--
			continue
		}

		b[lp], b[rp] = b[rp], b[lp]
		lp++
		rp--
	}

	return string(b)
}

// 是否A-Z a-z 0-9
func isAZ(char uint8) bool {
	if (char >= 97 && char <= 122) || (char >= 65 && char <= 90) {
		return true
	}
	return false
}
