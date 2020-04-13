package tree

//执行用时 :
//8 ms
//, 在所有 Go 提交中击败了
//79.31%
//的用户
//内存消耗 :
//4.2 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

//有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。
//
//给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，让你重新上色这幅图像。
//
//为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。
//
//最后返回经过上色渲染后的图像。
//
//示例 1:
//
//输入:
//image = [[1,1,1],[1,1,0],[1,0,1]]
//sr = 1, sc = 1, newColor = 2
//输出: [[2,2,2],[2,2,0],[2,0,1]]
//解析:
//在图像的正中间，(坐标(sr,sc)=(1,1)),
//在路径上所有符合条件的像素点的颜色都被更改成2。
//注意，右下角的像素没有更改为2，
//因为它不是在上下左右四个方向上与初始点相连的像素点。
//注意:
//
//image 和 image[0] 的长度在范围 [1, 50] 内。
//给出的初始点将满足 0 <= sr < image.length 和 0 <= sc < image[0].length。
//image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535]内。
//通过次数11,794提交次数22,151
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/flood-fill
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	height := len(image)
	//如果没有数据，直接返回image
	if height == 0 {
		return image
	}

	width := len(image[0])
	if width == 0 {
		return image
	}

	//如果目标坐标超出范围 返回image
	if sr > height-1 || sc > width-1 {
		return image
	}

	//要变成的颜色
	cor := image[sr][sc]

	//如果要替换的颜色和目标颜色相同，那就不要动了
	if cor == newColor {
		return image
	}

	var queue [][]int

	//要处理的都放在列队中
	queue = append(queue, []int{sr, sc})
	for len(queue) > 0 {
		tLen := len(queue)
		for i := 0; i < tLen; i++ {
			//根据坐标寻找坐标的上下左右，同时将当前坐标变成newColor
			yx := queue[i]
			y := yx[0]
			x := yx[1]
			image[y][x] = newColor

			//上
			if y-1 >= 0 && image[y-1][x] == cor {
				image[y-1][x] = newColor
				queue = append(queue, []int{y - 1, x})
			}

			//下
			if y+1 <= height-1 && image[y+1][x] == cor {
				image[y+1][x] = newColor
				queue = append(queue, []int{y + 1, x})
			}

			//左
			if x-1 >= 0 && image[y][x-1] == cor {
				image[y][x-1] = newColor
				queue = append(queue, []int{y, x - 1})
			}

			//右
			if x+1 <= width-1 && image[y][x+1] == cor {
				image[y][x+1] = newColor
				queue = append(queue, []int{y, x + 1})
			}
		}
		queue = queue[tLen:]
	}

	return image
}
