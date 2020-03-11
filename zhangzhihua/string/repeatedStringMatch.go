package string

import "fmt"

//给定两个字符串 A 和 B, 寻找重复叠加字符串A的最小次数，
// 使得字符串B成为叠加后的字符串A的子串，如果不存在则返回 -1。
//
//举个例子，A = "abcd"，B = "cdabcdab"。
//
//答案为 3， 因为 A 重复叠加三遍后为 “abcdabcdabcd”，
// 此时 B 是其子串；A 重复叠加两遍后为"abcdabcd"，B 并不是其子串。
//
//注意:
//
// A 与 B 字符串的长度在1和10000区间范围内。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/repeated-string-match
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

	//a := "abcd"
	//b := "cdabcdab"

	//a := "abc"
	//b := "cabcabca"

	a := "abcabcabcabc"
	b := "abac"

	//执行用时 :
	//	32 ms
	//	, 在所有 golang 提交中击败了
	//	32.14%
	//		的用户
	//内存消耗 :
	//	8.8 MB
	//	, 在所有 golang 提交中击败了
	//	25.00%
	//		的用户

	fmt.Println(repeatedStringMatch(a, b))
}

func repeatedStringMatch(A string, B string) int {

	BLen := len(B)
	ALen := len(A)

	cnt := BLen/ALen + 2

	cpA := A
	bl := false
	num := 1
	for i := 0; i <= cnt; i++ {

		bl = isSubStr(A, B)
		if bl {
			break
		}
		num++
		A += cpA
	}

	if !bl {
		return -1
	}
	return num
}

/**
subStr是否是str的字串
*/
func isSubStr(str string, subStr string) bool {
	strLen := len(str)
	subStrLen := len(subStr)

	if strLen == 0 || subStrLen == 0 || subStrLen > strLen {
		return false
	}

	//思路根据B字符串的的长度在str中查找
	for i := 0; i < strLen && subStrLen+i <= strLen; i++ {
		if str[i:i+subStrLen] == subStr {
			return true
		}
	}

	return false
}
