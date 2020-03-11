package string

import (
	"fmt"
	"strconv"
)

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
//注意：
//
//num1 和num2 的长度都小于 5100.
//num1 和num2 都只包含数字 0-9.
//num1 和num2 都不包含任何前导零。
//你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

	num1 := "0"
	num2 := "0"

	//性能不是忒别好

	//执行用时 :
	//8 ms
	//, 在所有 golang 提交中击败了
	//57.92%
	//的用户
	//内存消耗 :
	//7.3 MB
	//, 在所有 golang 提交中击败了
	//18.42%
	//的用户
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {

	n1Len := len(num1)
	n2Len := len(num2)

	i := n1Len - 1
	j := n2Len - 1

	str := ""
	var carry uint8 = 0
	var tmp1 uint8 = 0
	var tmp2 uint8 = 0
	for i >= 0 || j >= 0 {

		if i >= 0 {
			tmp1 = num1[i] - 48
		}

		if j >= 0 {
			tmp2 = num2[j] - 48
		}

		t := tmp1 + tmp2 + carry

		if t >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		str = strconv.Itoa(int(t%10)) + str

		tmp1 = 0
		tmp2 = 0
		i--
		j--
	}

	if carry == 1 {
		str = "1" + str
	}

	return str
}
