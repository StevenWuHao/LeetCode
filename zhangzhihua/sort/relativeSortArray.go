package sort

import "fmt"

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
