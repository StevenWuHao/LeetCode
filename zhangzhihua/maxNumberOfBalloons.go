package main

import (
	"fmt"
	"math"
)

//
//给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。
//
//字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。
//
//
//示例 1：
//
//输入：text = "nlaebolko"
//输出：1
//示例 2：
//
//
//输入：text = "loonbalxballpoon"
//输出：2
//示例 3：
//
//输入：text = "leetcode"
//输出：0
//
//
//提示：
//
//1 <= text.length <= 10^4
//text 全部由小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-number-of-balloons
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	text := "loonbalxballpoonoo"

	//执行用时 :
	//	0 ms
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户
	//内存消耗 :
	//	2.2 MB
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户

	fmt.Println(maxNumberOfBalloons(text))
}

func maxNumberOfBalloons(text string) int {
	ballons := make(map[byte]int)
	ballons['b'] = 1
	ballons['a'] = 1
	ballons['l'] = 2
	ballons['o'] = 2
	ballons['n'] = 1

	m := make(map[byte]int)

	i := len(text) - 1
	for i >= 0 {
		_, ok := ballons[text[i]]
		if ok {
			m[text[i]]++
		}
		i--
	}

	//如果长度不同直接返回
	if len(ballons) != len(m) {
		return 0
	}

	for k := range m {
		tmp := math.Floor(float64(m[k] / ballons[k]))
		ballons[k] = int(tmp)
	}

	num := -1
	//求最小值
	for _, v := range ballons {
		if num == -1 {
			num = v
		}
		if v < num {
			num = v
		}
	}

	if num == -1 {
		num = 0
	}
	return num
}
