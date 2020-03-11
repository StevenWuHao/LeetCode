package string

import (
	"fmt"
	"strings"
)

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//示例 1:
//
//输入: "Let's take LeetCode contest"
//输出: "s'teL ekat edoCteeL tsetnoc"
//注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

	//b := []byte{'a', 'b', 'c'}
	//fmt.Println(reverse(b))
	//
	//str := "Let's take LeetCode contest"
	str := "I love u"
	fmt.Println(reverseWords(str))
}

func reverseWords(s string) string {

	sLen := len(s)
	if sLen == 0 || sLen == 1 {
		return s
	}
	fp := 0
	lp := 0
	var str strings.Builder
	for i := 0; i < sLen; i++ {

		if s[i] == 32 || i == sLen-1 {
			if i == sLen-1 {
				str.WriteString(reverse([]byte(s[fp:sLen])))
			} else {
				str.WriteString(reverse([]byte(s[fp : lp+1])))
				str.WriteString(" ")
			}

			fp = i + 1
			lp = fp
		}
		lp = i
	}

	return str.String()
}

func reverse(b []byte) string {

	bLen := len(b)
	if bLen == 0 || bLen == 1 {
		return string(b)
	}
	i := 0
	j := bLen - 1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
