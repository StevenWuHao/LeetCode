package tree

import (
	"fmt"
	"testing"
)

func TestOrangesRotting(t *testing.T) {

	//time := orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
	//time := orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}})
	time := orangesRotting([][]int{{1}})
	fmt.Printf("需要%d秒,苹果都烂了 \r\n", time)
}
