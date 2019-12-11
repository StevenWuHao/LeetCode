package main

import (
	"fmt"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "){"
	//s := ""
	//s := "()"
	//s := "()[]{}"
	//s := "(]"
	//s := "([)]"
	//s := "{[]}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {

	sLen := len(s)

	if sLen == 0 {
		return true
	}

	//如果不是偶数，直接返回
	if sLen%2 != 0 {
		return false
	}

	//asc对照表
	m := make(map[uint8]uint8)
	m[40] = 41   //(=72  )=73
	m[123] = 125 //{=123 }=125
	m[91] = 93   //[=91  ]=93

	stack := make([]uint8, 0)

	for i := 0; i < sLen; i++ {
		_, ok := m[s[i]]

		if ok {
			//如果左侧的括号，放到栈中，等待抵消
			stack = append(stack, s[i])
			continue
		}
		//如果不是开始的括号，判断栈的最后一位，是否和当前的能抵消
		//如果len(stack) == 0  表示栈中可能要出现右侧的括号，应该直接false
		//可以抵消，将栈的最后一位移除
		if len(stack) == 0 || m[stack[len(stack)-1]] != s[i] {
			return false
		}
		stack = stack[:len(stack)-1]
		continue
	}

	//最后栈中应该是空的
	stackLen := len(stack)

	if stackLen != 0 {
		return false
	}

	return true
}
