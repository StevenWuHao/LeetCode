package tree

import "fmt"

//执行用时 :
//12 ms
//, 在所有 Go 提交中击败了
//12.66%
//的用户
//内存消耗 :
//3.2 MB
//, 在所有 Go 提交中击败了
//40.00%
//的用户

/**
994. 腐烂的橘子
在给定的网格中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。



示例 1：



输入：[[2,1,1],[1,1,0],[0,1,1]]
输出：4
示例 2：

输入：[[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
示例 3：

输入：[[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。


提示：

1 <= grid.length <= 10
1 <= grid[0].length <= 10
grid[i][j] 仅为 0、1 或 2
*/
func orangesRotting(grid [][]int) int {

	time := 0
	var queue [][]int //存放坏掉的橘子
	goodOrg := 0

	//首先找到滥掉的橘子，把烂橘子放到列队里面去
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 2 {
				fmt.Printf("烂苹果的位置:y:%d x:%d \n\r", y, x)
				queue = append(queue, []int{y, x})
			} else if grid[y][x] == 1 {
				fmt.Printf("美丽苹果的位置:y:%d x:%d \n\r", y, x)
				goodOrg++
			}
		}
	}

	fmt.Printf("坏苹果总数： %d 个 好苹果总数: %d 个 \n\r ", len(queue), goodOrg)

	h := len(grid) - 1    //格子的高度
	w := len(grid[0]) - 1 //格子的宽度

	//循环找烂橘子，将烂橘子周围上下左右变成烂橘子
	for len(queue) > 0 && goodOrg > 0 {
		time++
		fmt.Printf("time:%d \r\n ========================= \r\n", time)
		tLen := len(queue)
		for i := 0; i < tLen; i++ {
			coordinate := queue[i]
			y := coordinate[0]
			x := coordinate[1]

			//上
			if y-1 >= 0 && grid[y-1][x] == 1 {
				fmt.Printf("坐标：y:%d x:%d 上 是好苹果 坐标 y:%d x:%d 当前苹果值：%d \n\r", y, x, y-1, x, grid[y-1][x])
				grid[y-1][x] = 2
				queue = append(queue, []int{y - 1, x})
				goodOrg--
			}

			//下
			if y+1 <= h && grid[y+1][x] == 1 {
				fmt.Printf("坐标：y:%d x:%d 下 是好苹果 坐标 y:%d x:%d  当前苹果值：%d \n\r", y, x, y+1, x, grid[y+1][x])
				grid[y+1][x] = 2
				queue = append(queue, []int{y + 1, x})
				goodOrg--
			}

			//左
			if x-1 >= 0 && grid[y][x-1] == 1 {
				fmt.Printf("坐标：y:%d x:%d 左边 是好苹果 坐标 y:%d x:%d  当前苹果值：%d \n\r", y, x, y, x-1, grid[y][x-1])
				grid[y][x-1] = 2
				queue = append(queue, []int{y, x - 1})
				goodOrg--
			}

			//右
			if x+1 <= w && grid[y][x+1] == 1 {
				fmt.Printf("坐标：y:%d x:%d 右 是好苹果 坐标 y:%d x:%d  当前苹果值：%d \n\r", y, x, y, x+1, grid[y][x+1])
				grid[y][x+1] = 2
				queue = append(queue, []int{y, x + 1})
				goodOrg--
			}
		}
		queue = queue[tLen:]
	}
	if goodOrg > 0 { //还有健康的
		return -1
	} else {
		return time
	}
}
