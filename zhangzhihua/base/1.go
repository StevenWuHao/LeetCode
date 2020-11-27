package main

import (
	"fmt"
	"sync"
)

/**
锁
读锁
获取到读锁之后，再获取同一个读锁程序会不会阻塞
结论不会阻塞
*/
func main() {
	var rw sync.RWMutex

	fmt.Println("准备获取读锁")
	rw.RLock()
	fmt.Println("已经获取到读锁")

	fmt.Println("再获取一个读锁")
	rw.RLock()
	fmt.Println("如果获取了，说明读锁不阻塞")

}
