package main

import (
	"fmt"
	"strings"
)

//给定一个段落 (paragraph) 和一个禁用单词列表 (banned)。返回出现次数最多，同时不在禁用列表中的单词。题目保证至少有一个词不在禁用列表中，而且答案唯一。
//
//禁用列表中的单词用小写字母表示，不含标点符号。段落中的单词不区分大小写。答案都是小写字母。
//
//
//
//示例：
//
//输入:
//paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
//banned = ["hit"]
//输出: "ball"
//解释:
//"hit" 出现了3次，但它是一个禁用的单词。
//"ball" 出现了2次 (同时没有其他单词出现2次)，所以它是段落里出现次数最多的，且不在禁用列表中的单词。
//注意，所有这些单词在段落里不区分大小写，标点符号需要忽略（即使是紧挨着单词也忽略， 比如 "ball,"），
//"hit"不是最终的答案，虽然它出现次数更多，但它在禁用单词列表中。
//
//
//说明：
//
//1 <= 段落长度 <= 1000.
//1 <= 禁用单词个数 <= 100.
//1 <= 禁用单词长度 <= 10.
//答案是唯一的, 且都是小写字母 (即使在 paragraph 里是大写的，即使是一些特定的名词，答案都是小写的。)
//paragraph 只包含字母、空格和下列标点符号!?',;.
//不存在没有连字符或者带有连字符的单词。
//单词里只包含字母，不会出现省略号或者其他标点符号。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/most-common-word
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	//paragraph := "a."
	banned := []string{"hit"}

	//执行用时 :
	//	0 ms
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户
	//内存消耗 :
	//	2.4 MB
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户
	fmt.Println(mostCommonWord(paragraph, banned))
}

func mostCommonWord(paragraph string, banned []string) string {

	m := make(map[string]int)
	cnt := 0
	ret := ""

	//全部转成小写
	paragraph = strings.ToLower(paragraph)
	paragraphLen := len(paragraph)

	firstWordsIndex := -1
	for read := 0; read < paragraphLen; read++ {

		//遇到非字母&值钱是单词
		if !isAZx(paragraph[read]) || read == paragraphLen-1 {
			//检查之前是否有字母需要处理
			if firstWordsIndex == -1 {
				continue
			}

			//排除单词
			if inBanned(paragraph[firstWordsIndex:read], banned) {
				firstWordsIndex = -1
				continue
			}

			//最后一位处理
			if read == paragraphLen-1 {
				//检查最后一位是否是字母
				if !isAZx(paragraph[read]) {
					m[paragraph[firstWordsIndex:paragraphLen-1]]++
				} else {
					m[paragraph[firstWordsIndex:paragraphLen]]++
				}
				firstWordsIndex = -1
				continue
			}

			m[paragraph[firstWordsIndex:read]]++
			firstWordsIndex = -1
			continue
		}

		//遇到非字母
		if firstWordsIndex == -1 {
			firstWordsIndex = read
		}
	} //for

	fmt.Println(m)
	for k, v := range m {
		if v > cnt {
			cnt = v
			ret = k
		}
	}

	return ret
}

// 是否A-Z a-z 0-9
func isAZx(char uint8) bool {
	if (char >= 97 && char <= 122) || (char >= 65 && char <= 90) {
		return true
	}
	return false
}

func inBanned(s string, banned []string) bool {
	bl := false
	for _, v := range banned {
		if v == s {
			bl = true
			break
		}
	}

	return bl
}
