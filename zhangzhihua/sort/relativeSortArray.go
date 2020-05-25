package sort

import "fmt"

//1122. 数组的相对排序
//给你两个数组，arr1 和 arr2，
//
//arr2 中的元素各不相同
//arr2 中的每个元素都出现在 arr1 中
//对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
//
//
//
//示例：
//
//输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
//输出：[2,2,2,1,4,3,3,9,6,7,19]
//
//
//提示：
//
//arr1.length, arr2.length <= 1000
//0 <= arr1[i], arr2[i] <= 1000
//arr2 中的元素 arr2[i] 各不相同
//arr2 中的每个元素 arr2[i] 都出现在 arr1 中

//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2.5 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户
func relativeSortArray(arr1 []int, arr2 []int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(arr2); i++ {
		m[arr2[i]] = 0
	}

	//找出不在arr2中的数据,在的话累计下，后面用
	var notInArr2 []int
	for i := 0; i < len(arr1); i++ {
		if _, ok := m[arr1[i]]; !ok {
			notInArr2 = append(notInArr2, arr1[i])
		} else {
			m[arr1[i]]++
		}
	}
	fmt.Println(notInArr2)
	//将不再arr2中的排序
	for i := 0; i < len(notInArr2); i++ {
		for j := i + 1; j < len(notInArr2); j++ {
			if notInArr2[i] > notInArr2[j] {
				notInArr2[i], notInArr2[j] = notInArr2[j], notInArr2[i]
			}
		}
	}
	var ret []int
	for i := 0; i < len(arr2); i++ {
		ret = append(ret, arr2[i])
		cnt, _ := m[arr2[i]]
		for cnt-1 > 0 {
			ret = append(ret, arr2[i])
			cnt--
		}
	}

	return append(ret, notInArr2...)
}
