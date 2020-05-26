package sort

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
