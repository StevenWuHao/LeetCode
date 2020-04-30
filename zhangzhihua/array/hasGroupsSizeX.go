package array

//执行用时 :
//16 ms
//, 在所有 Go 提交中击败了
//93.47%
//的用户
//内存消耗 :
//6 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//给定一副牌，每张牌上都写着一个整数。
//
//此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
//
//每组都有 X 张牌。
//组内所有的牌上都写着相同的整数。
//仅当你可选的 X >= 2 时返回 true。
//
//
//
//示例 1：
//
//输入：[1,2,3,4,4,3,2,1]
//输出：true
//解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
//示例 2：
//
//输入：[1,1,1,2,2,2,3,3]
//输出：false
//解释：没有满足要求的分组。
//示例 3：
//
//输入：[1]
//输出：false
//解释：没有满足要求的分组。
//示例 4：
//
//输入：[1,1]
//输出：true
//解释：可行的分组是 [1,1]
//示例 5：
//
//输入：[1,1,2,2,2,2]
//输出：true
//解释：可行的分组是 [1,1]，[2,2]，[2,2]
//
//提示：
//
//1 <= deck.length <= 10000
//0 <= deck[i] < 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func hasGroupsSizeX(deck []int) bool {

	deckLen := len(deck)
	if deckLen < 2 {
		return false
	}

	//先分组，统计每个数的出现次数
	groupCntMap := make(map[int]int)
	for i := 0; i < len(deck); i++ {
		if m, ok := groupCntMap[deck[i]]; ok {
			groupCntMap[deck[i]] = m + 1
		} else {
			groupCntMap[deck[i]] = 1
		}
	}

	g := -1
	for _, c := range groupCntMap {
		if g == -1 {
			g = c
		} else {
			if g > c {
				if gcd(g, c) < 2 {
					return false
				}
			} else {
				if gcd(c, g) < 2 {
					return false
				}
			}
		}
	}
	return true
}

func gcd(x int, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
