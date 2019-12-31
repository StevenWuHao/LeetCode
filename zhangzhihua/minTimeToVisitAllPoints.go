package main

import (
	"fmt"
	"math"
)

//平面上有 n 个点，点的位置用整数坐标表示 points[i] = [xi, yi]。请你计算访问所有这些点需要的最小时间（以秒为单位）。
//
//你可以按照下面的规则在平面上移动：
//
//每一秒沿水平或者竖直方向移动一个单位长度，或者跨过对角线（可以看作在一秒内向水平和竖直方向各移动一个单位长度）。
//必须按照数组中出现的顺序来访问这些点。
//
//
//示例 1：
//
//
//
//输入：points = [[1,1],[3,4],[-1,0]]
//输出：7
//解释：一条最佳的访问路径是： [1,1] -> [2,2] -> [3,3] -> [3,4] -> [2,3] -> [1,2] -> [0,1] -> [-1,0]
//从 [1,1] 到 [3,4] 需要 3 秒
//从 [3,4] 到 [-1,0] 需要 4 秒
//一共需要 7 秒
//示例 2：
//
//输入：points = [[3,2],[-2,2]]
//输出：5
//
//
//提示：
//
//points.length == n
//1 <= n <= 100
//points[i].length == 2
//-1000 <= points[i][0], points[i][1] <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-time-visiting-all-points
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	//执行用时 :
	//	4 ms
	//	, 在所有 golang 提交中击败了
	//	92.86%
	//		的用户
	//内存消耗 :
	//	3.4 MB
	//	, 在所有 golang 提交中击败了
	//	100.00%
	//		的用户
	fmt.Println(minTimeToVisitAllPoints(points))
}

func minTimeToVisitAllPoints(points [][]int) int {
	x, y := points[0][0], points[0][1]
	miao := 0.0
	for i := 1; i < len(points); i++ {
		tmpX := points[i][0]
		tmpY := points[i][1]
		miao += math.Max(math.Abs(float64(tmpX-x)), math.Abs(float64(tmpY-y)))
		x = tmpX
		y = tmpY
	}

	return int(miao)
}
