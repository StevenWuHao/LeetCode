package main

import "fmt"

//给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回 true ；否则返回 false。
//
//(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。)
//
//注意：
//
//你可以假设两个字符串均只含有小写字母。
//
//canConstruct("a", "b") -> false
//canConstruct("aa", "ab") -> false
//canConstruct("aa", "aab") -> true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/ransom-note
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

	//没有完成，因为下面在leetcode返回竟然是true，完全不懂
	ransomNote := ""
	magazine := "a"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {

	rLen := len(ransomNote)
	mLen := len(magazine)

	//如果相等直接返回
	if ransomNote == magazine {
		return true
	}

	if rLen == mLen {
		return true
	}

	cnt := 0
	i := 0
	for i < mLen {
		j := 0
		for j < rLen {
			if magazine[i] == ransomNote[j] {
				cnt++
				if cnt == rLen {
					return true
				}
				break
			}
			j++
		}
		i++
	}

	return false
}
