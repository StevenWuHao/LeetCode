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
	s := "a a abn "
	//s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	sLen := len(s)

	if sLen == 0 {
		return 0
	}

	nullLastIndex := -1 //空字符串最后出现所在的位置
	AToZLastIndex := -1 //A-Z最后出现的位置,-1代表没有
	cnt := 0
	AToZLastMaxLen := 0
	for k, v := range s {
		if v == 32 {
			nullLastIndex = k
			cnt = 0
			continue
		}

		//非空
		if cnt == 0 {
			AToZLastIndex = k
			AToZLastMaxLen = 0
		}

		cnt++
		if cnt > AToZLastMaxLen {
			AToZLastMaxLen = cnt
		}
	}

	//全是空，没有单个字符直接返回0
	if AToZLastIndex == -1 {
		return 0
	}

	//没有空字符串，直接返回长度
	if nullLastIndex == -1 {
		return sLen
	}

	return len(s[AToZLastIndex : AToZLastIndex+AToZLastMaxLen])
}
