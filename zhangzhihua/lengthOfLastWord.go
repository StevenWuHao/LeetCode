package main

import "fmt"

//给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
//
//如果不存在最后一个单词，请返回 0 。
//
//说明：一个单词是指由字母组成，但不包含任何空格的字符串。
//
//示例:
//
//输入: "Hello World"
//输出: 5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/length-of-last-word
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "a a abn sdfsdfsfddsf sdfsdf sadfsdfs fdsfasdfasfdafsafdafdfds dfsa asd fasf asf sa fsdf sdfsadf sdf sdf as fas fs fasf "
	//s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {

	sLen := len(s)

	if sLen == 0 {
		return 0
	}

	//空字符串最后出现所在的位置,如果！=-1说明整个字符串没有空字符，可以直接返回字符串的长度
	nullLastIndex := -1

	wordsStartIndex := -1 //最后单词的位置
	wordsMaxLen := 0      //最后单词的长度
	isNewWords := true    //是否新单词开始

	for k, v := range s {

		if v == 32 {
			nullLastIndex = k
			isNewWords = true //空字符出现，下一个一定是新单词
			continue
		}
		//到这这里说明是非空格的字符

		//新单词
		if isNewWords {
			wordsStartIndex = k //记录位置起始位置
			wordsMaxLen = 0     //重置 最后单词的长度
		}

		isNewWords = false
		wordsMaxLen++
	}

	//全是空，没有单个字符直接返回0
	if wordsStartIndex == -1 {
		return 0
	}

	//没有空字符串，直接返回长度
	if nullLastIndex == -1 {
		return sLen
	}
	return len(s[wordsStartIndex : wordsStartIndex+wordsMaxLen])
}
