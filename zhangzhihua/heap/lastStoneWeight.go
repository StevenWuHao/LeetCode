package heap

//1046. 最后一块石头的重量
//有一堆石头，每块石头的重量都是正整数。
//
//每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
//
//如果 x == y，那么两块石头都会被完全粉碎；
//如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
//最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。
//
//
//
//示例：
//
//输入：[2,7,4,1,8,1]
//输出：1
//解释：
//先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
//再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
//接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
//最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。

//12 分钟前	通过	0 ms	2 MB	Golang
func lastStoneWeight(stones []int) int {
	if len(stones) < 3 {
		for i := 3 - len(stones) - 1; i >= 0; i-- {
			stones = append(stones, 0)
		}
	}
	heapSort(stones)
	for stones[0] != 0 && stones[1] != 0 {
		if stones[0] == stones[1] {
			stones[0], stones[1] = 0, 0
		} else if stones[0] != stones[1] {
			stones[0] = stones[0] - stones[1]
			stones[1] = 0
		}
		heapSort(stones)
	}
	return stones[0]
}

func heapSort(nums []int) {
	buildMaxHeap(nums)
	for j := len(nums) - 1; j > 0; j-- {
		swap(0, j, nums)
		heapify(0, j, nums)
	}
}

func buildMaxHeap(nums []int) {
	l := len(nums)
	for i := l/2 - 1; i >= 0; i-- {
		heapify(i, l, nums)
	}
}

func heapify(idx, length int, nums []int) {

	left := 2*idx + 1
	right := 2*idx + 2
	largest := idx
	if left < length && nums[left] < nums[idx] {
		largest = left
	}

	if right < length && nums[right] < nums[largest] {
		largest = right
	}

	if largest != idx {
		swap(largest, idx, nums)
		heapify(largest, length, nums)
	}
}

func swap(i, j int, nums []int) {
	nums[i], nums[j] = nums[j], nums[i]
}
