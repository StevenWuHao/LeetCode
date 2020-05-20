package sort

//922. 按奇偶排序数组 II
//给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。
//
//对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。
//
//你可以返回任何满足上述条件的数组作为答案。
//
//
//
//示例：
//
//输入：[4,2,5,7]
//输出：[4,5,2,7]
//解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
//
//
//提示：
//
//2 <= A.length <= 20000
//A.length % 2 == 0
//0 <= A[i] <= 1000

//执行用时 :
//28 ms
//, 在所有 Go 提交中击败了
//71.92%
//的用户
//内存消耗 :
//6.3 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户
func sortArrayByParityII(A []int) []int {

	for i := 0; i < len(A); i++ {
		isEventI := isEvenNumbers(i)
		isEventAI := isEvenNumbers(A[i])
		if isEventI && isEventAI || !isEventI && !isEventAI {
			continue
		}
		for j := i + 1; j < len(A); j++ {
			if isEventI && isEvenNumbers(A[j]) { //索引偶数
				A[i], A[j] = A[j], A[i]
				break
			}
			if !isEventI && !isEvenNumbers(A[j]) { //索引奇数
				A[i], A[j] = A[j], A[i]
				break
			}
		}
	}
	return A
}

func isEvenNumbers(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
