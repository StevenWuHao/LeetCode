package sort

func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	}

	if s == t {
		return true
	}

	alpha := [26]int{}
	for i := 0; i < len(s); i++ {
		alpha[s[i]]++
		alpha[t[i]]--
	}
	for i := 0; i < len(alpha); i++ {
		if alpha[i] != 0 {
			return false
		}
	}
	return true
}

/**
方法2 哈希表
*/
func isAnagramH(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	}

	if s == t {
		return true
	}

	hash := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			hash[s[i]]++
		} else {
			hash[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := hash[t[i]]; !ok {
			return false
		} else {
			if hash[t[i]] == 0 {
				return false
			}
			hash[t[i]]--
		}
	}

	return true
}

/**
方法1：冒泡排序，超时
*/
func isAnagram1(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	}
	if sort(s) == sort(t) {
		return true
	}
	return false
}

func sort(s string) string {
	newS := make([]uint8, len(s))
	for i := 0; i < len(s); i++ {
		newS = append(newS, s[i])
	}
	for i := 0; i < len(newS); i++ {
		for j := i + 1; j < len(newS); j++ {
			if newS[i] > newS[j] {
				newS[i], newS[j] = newS[j], newS[i]
			}
		}
	}
	return string(newS)
}
