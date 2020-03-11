package string

import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-common-prefix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

	//strs := []string{"flower", "flow", "flight", "fly"}
	//strs := []string{"flower", "flower", "flower", "flewer"}
	strs := []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {

	strsLen := len(strs)

	//如果传入空数组，直接返回空字符串
	if strsLen == 0 {
		return ""
	}

	//如果传入的数据长度=1，则直接返回这个值的第一个值
	if strsLen == 1 {
		return strs[0]
	}
	//把数组的第一位作为模版
	first := strs[0]

	for f := 0; f < len(first); f++ {
		//从数组第二位开始
		for i := 1; i < strsLen; i++ {
			//如果和后面一位不同，说明直接可以返回了返回长度就是模版的F字符串位置
			if f == len(strs[i]) || first[f] != strs[i][f] {
				return first[0:f]
			}
		}
	}

	return first
}
