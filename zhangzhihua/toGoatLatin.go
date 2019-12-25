package main

import (
	"fmt"
	"strings"
)

//给定一个由空格分割单词的句子 S。每个单词只包含大写或小写字母。
//
//我们要将句子转换为 “Goat Latin”（一种类似于 猪拉丁文 - Pig Latin 的虚构语言）。
//
//山羊拉丁文的规则如下：
//
//如果单词以元音开头（a, e, i, o, u），在单词后添加"ma"。
//例如，单词"apple"变为"applema"。
//
//如果单词以辅音字母开头（即非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。
//例如，单词"goat"变为"oatgma"。
//
//根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从1开始。
//例如，在第一个单词后添加"a"，在第二个单词后添加"aa"，以此类推。
//返回将 S 转换为山羊拉丁文后的句子。
//
//示例 1:
//
//输入: "I speak Goat Latin"
//输出: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
//示例 2:
//
//输入: "The quick brown fox jumped over the lazy dog"
//输出: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
//说明:
//
//S 中仅包含大小写字母和空格。单词间有且仅有一个空格。
//1 <= S.length <= 150。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/goat-latin
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

	s := "I speak Goat Latin"
	fmt.Println(toGoatLatin(s))
}

func toGoatLatin(S string) string {

	sLen := len(S)

	isVowel := false
	firstWord := -1
	str := strings.Builder{}
	wordNum := 1

	for read := 0; read < sLen; read++ {

		//不为空，记录首字母出现的位置,同时检查检查首字母是否元音
		if S[read] != 32 {
			if firstWord == -1 {
				firstWord = read
				isVowel = isVowelsX(S[read])
			}
			continue
		}

		//如果为空,判断首字母是否元音， 如果是，在单词末尾添加ma。如果不是，
		// 移除第一个字符并将它放到末尾，之后再添加"ma"。最后添加索引位a
		if isVowel {
			str.Write([]byte(S[firstWord:read]))
		} else {
			str.Write([]byte(S[firstWord+1 : read]))
			str.Write([]byte(S[firstWord : firstWord+1]))
		}
		str.Write([]byte{'m', 'a'})

		j := wordNum
		for j > 0 {
			str.Write([]byte{'a'})
			j--
		}

		wordNum++
		str.Write([]byte{' '})
		firstWord = -1
		isVowel = false
	}

	//这里有个边界，就是处理最后一个单词
	if firstWord != -1 {
		if isVowel {
			str.Write([]byte(S[firstWord:sLen]))
		} else {
			str.Write([]byte(S[firstWord+1 : sLen]))
			str.Write([]byte(S[firstWord : firstWord+1]))
		}
		str.Write([]byte{'m', 'a'})
		for wordNum > 0 {
			str.Write([]byte{'a'})
			wordNum--
		}
	}

	return str.String()
}
func isVowelsX(v uint8) bool {
	if v == 65 || v == 97 || v == 69 || v == 101 || v == 73 || v == 105 || v == 79 || v == 111 || v == 85 || v == 117 {
		return true
	}
	return false
}
