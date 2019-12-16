package main

import "fmt"

//编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
//
//示例 1:
//
//输入: "hello"
//输出: "holle"
//示例 2:
//
//输入: "leetcode"
//输出: "leotcede"
//说明:
//元音字母 a e i o u
//元音字母不包含字母"y"。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-vowels-of-a-string
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	str := "leetcode"
	fmt.Println(reverseVowels(str))
}

func reverseVowels(s string) string {
	p := []byte(s)

	lp := 0
	rp := len(p) - 1

	for lp < rp {

		lpB := isVowels(p[lp])
		rpB := isVowels(p[rp])

		if lpB && rpB {
			p[lp], p[rp] = p[rp], p[lp]
		}

		if !lpB {
			lp++
			continue
		}

		if !rpB {
			rp--
			continue
		}

		lp++
		rp--
	}

	return string(p)
}

func isVowels(v uint8) bool {
	if v == 65 || v == 97 || v == 69 || v == 101 || v == 73 || v == 105 || v == 79 || v == 111 || v == 85 || v == 117 {
		return true
	}
	return false
}
