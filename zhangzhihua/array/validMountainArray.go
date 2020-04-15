package array

//执行用时 :
//28 ms
//, 在所有 Go 提交中击败了
//97.78%
//的用户
//内存消耗 :
//6.2 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//941. 有效的山脉数组
//给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。
//
//让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：
//
//A.length >= 3
//在 0 < i < A.length - 1 条件下，存在 i 使得：
//A[0] < A[1] < ... A[i-1] < A[i]
//A[i] > A[i+1] > ... > A[A.length - 1]
//
//
//
//
//
//
//示例 1：
//
//输入：[2,1]
//输出：false
//示例 2：
//
//输入：[3,5,5]
//输出：false
//示例 3：
//
//输入：[0,3,2,1]
//输出：true
//
//
//提示：
//
//0 <= A.length <= 10000
//0 <= A[i] <= 10000
//

func validMountainArray(A []int) bool {

	ALen := len(A)
	if ALen < 3 {
		return false
	}

	//是否上坡
	isUp := false
	if A[0] < A[1] {
		isUp = true
	} else {
		return false
	}

	for i := 1; i < ALen && i+1 < ALen; i++ {
		if isUp {
			//平地，返回false
			if A[i+1] == A[i] {
				return false
			}
			//下坡了
			if A[i+1] < A[i] {
				isUp = false
				continue
			}
			continue
		}
		//下坡
		if A[i+1] >= A[i] {
			return false
		}
	}

	if isUp == true {
		return false
	}

	return true
}
