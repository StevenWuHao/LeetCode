package string

import (
	"fmt"
	"strings"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-palindrome
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str))
}

func isPalindrome(s string) bool {

	sLen := len(s)

	s = strings.ToLower(s)

	lp := 0
	rp := sLen - 1
	for lp < rp {
		if !isAZAnd19(s[lp]) {
			lp++
			continue
		}

		if !isAZAnd19(s[rp]) {
			rp--
			continue
		}

		if s[lp] != s[rp] {
			return false
		}
		lp++
		rp--
	}

	return true
}

// 是否A-Z a-z 0-9
func isAZAnd19(char uint8) bool {
	if (char >= 48 && char <= 57) || (char >= 97 && char <= 122) || (char >= 65 && char <= 90) {
		return true
	}
	return false
}
