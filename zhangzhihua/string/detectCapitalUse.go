package string

import "fmt"

//给定一个单词，你需要判断单词的大写使用是否正确。
//
//我们定义，在以下情况时，单词的大写用法是正确的：
//
//全部字母都是大写，比如"USA"。
//单词中所有字母都不是大写，比如"leetcode"。
//如果单词不只含有一个字母，只有首字母大写， 比如 "Google"。
//否则，我们定义这个单词没有正确使用大写字母。
//
//示例 1:
//
//输入: "USA"
//输出: True
//示例 2:
//
//输入: "FlaG"
//输出: False
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/detect-capital
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

	//word := "FlaG"
	//word := "USA"
	//word := "leetcode"
	//word := "Leetcode"
	//word := "LeetcodeX"
	word := "FFFFFFFFFFFFFFFFFFFFf"

	//执行用时 :
	//	0 ms
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户
	//内存消耗 :
	//	2 MB
	//	, 在所有 golang 提交中击败了
	//	82.35%
	//		的用户
	fmt.Println(detectCapitalUse(word))

}

func detectCapitalUse(word string) bool {

	wordLen := len(word)

	if wordLen == 0 {
		return false
	}

	if wordLen == 1 {
		return true
	}

	bl := true
	isFirstWordBig := false //首字母是否大写
	isAllBig := false       //是否必须都大写
	for i := 0; i < wordLen; i++ {

		isBig := false
		if word[i] >= 65 && word[i] <= 90 {
			isBig = true
		}

		//检查首字母是否大写,如果小写后面必须都是小写
		if i == 0 {
			if isBig {
				isFirstWordBig = true
			}
			continue
		}

		if i == 1 {
			if isFirstWordBig {
				if isBig {
					isAllBig = true
				}
			}
		}

		//如果首字母是小写，后面出现大写直接返回
		if isFirstWordBig == false {
			if isBig {
				bl = false
				break
			}
		}

		//如果必须全部大写
		if isAllBig {
			if !isBig {
				bl = false
				break
			}
		} else {
			if isBig {
				bl = false
				break
			}
		}
	}

	return bl
}
