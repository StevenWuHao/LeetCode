package string

import (
	"fmt"
	"strconv"
)

//给定两个二进制字符串，返回他们的和（用二进制表示）。
//
//输入为非空字符串且只包含数字 1 和 0。
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-binary
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	a := "1111"
	b := "1111"

	//a := "1010"
	//b := "1011"

	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {

	maxLen := 0
	aLen := len(a)
	bLen := len(b)

	//取一个长的二进制字符串
	if aLen >= bLen {
		maxLen = aLen
	} else {
		maxLen = bLen
	}

	//如果是空字符串直接返回字符串0
	if maxLen == 0 {
		return "0"
	}

	//将短的一位补0
	if aLen > bLen {
		for i := 0; i < maxLen-bLen; i++ {
			b = "0" + b
		}
	}

	if aLen < bLen {
		for i := 0; i < maxLen-aLen; i++ {
			a = "0" + a
		}
	}

	carry := 0 //进位
	str := ""
	maxLen = maxLen - 1
	for maxLen >= 0 {
		tmpA, _ := strconv.Atoi(string(a[maxLen]))
		tmpB, _ := strconv.Atoi(string(b[maxLen]))
		k := tmpA + tmpB + carry
		if k >= 2 {
			carry = 1
		} else {
			carry = 0
		}
		str = strconv.Itoa(k%2) + str
		maxLen--
	}

	if carry == 1 {
		str = strconv.Itoa(carry) + str
	}

	return str
}
