package main

import "fmt"

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//案例:
//
//s = "leetcode"
//返回 0.
//
//s = "loveleetcode",
//返回 2.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	s := "dfsdfsdfk"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {

	m0 := map[rune]int{}
	for _, v := range s {
		m0[v]++
	}
	for i, v := range s {
		if m0[v] == 1 {
			return i
		}
	}
	return -1
}
