package main

import (
	"fmt"
	"strconv"
)

//给定一组字符，使用原地算法将其压缩。
////
////压缩后的长度必须始终小于或等于原数组长度。
////
////数组的每个元素应该是长度为1 的字符（不是 int 整数类型）。
////
////在完成原地修改输入数组后，返回数组的新长度。
////
////
////
////进阶：
////你能否仅使用O(1) 空间解决问题？
////
////
////
////示例 1：
////
////输入：
////["a","a","b","b","c","c","c"]
////
////输出：
////返回6，输入数组的前6个字符应该是：["a","2","b","2","c","3"]
////
//说明：
//"aa"被"a2"替代。"bb"被"b2"替代。"ccc"被"c3"替代。
//示例 2：
//
//输入：
//["a"]
//
//输出：
//返回1，输入数组的前1个字符应该是：["a"]
//
//说明：
//没有任何字符串被替代。
//示例 3：
//
//输入：
//["a","b","b","b","b","b","b","b","b","b","b","b","b"]
//
//输出：
//返回4，输入数组的前4个字符应该是：["a","b","1","2"]。
//
//说明：
//由于字符"a"不重复，所以不会被压缩。"bbbbbbbbbbbb"被“b12”替代。
//注意每个数字在数组中都有它自己的位置。
//注意：
//
//所有字符都有一个ASCII值在[35, 126]区间内。
//1 <= len(chars) <= 1000。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-compression
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

	c := []byte{'a', 'a', 'a', 'c', 'c', 'c'}
	//c := []byte{'c'}
	//c := []byte{'a', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}
	//fmt.Printf("%s", c)

	//执行用时 :
	//	4 ms
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户
	//内存消耗 :
	//	2.9 MB
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户

	fmt.Println(compress(c))
}

func compress(chars []byte) int {
	cLen := len(chars)

	if cLen == 0 || cLen == 1 {
		return cLen
	}

	read, fp, fpTime := 0, 0, 0

	b := chars[0:0]

	for read <= cLen-1 {
		if read+1 == cLen || chars[read] != chars[read+1] {
			b = append(b, chars[fp:fp+1]...)
			if fpTime > 1 {
				b = append(b, []byte(strconv.Itoa(fpTime))...)
			}
			fp = read + 1
			fpTime = 0
		} else {
			if fpTime == 0 {
				fp = read
				fpTime = 2
			} else {
				fpTime++
			}
		}
		read++
	}

	return len(b)
}
