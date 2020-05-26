package sort

//976. 三角形的最大周长
//给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。
//
//如果不能形成任何面积不为零的三角形，返回 0。
//
//
//
//示例 1：
//
//输入：[2,1,2]
//输出：5
//示例 2：
//
//输入：[1,2,1]
//输出：0
//示例 3：
//
//输入：[3,2,3,4]
//输出：10
//示例 4：
//
//输入：[3,6,2,3]
//输出：8
//
//
//提示：
//
//3 <= A.length <= 10000
//1 <= A[i] <= 10^6

//2 分钟前	通过	1056 ms	6.2 MB	Golang

/**
方法1，先排序，然后从末端开始找
*/
func largestPerimeter(A []int) int {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
	lp := len(A) - 3
	rp := len(A)
	for lp >= 0 {
		if sum, ok := isTriangle(A[lp:rp]); ok {
			return sum
		}
		rp--
		lp--
	}
	return 0
}

func isTriangle(ary []int) (int, bool) {
	if ary[0]+ary[1] > ary[2] {
		return ary[0] + ary[1] + ary[2], true
	}
	return 0, false
}
