package main

import (
	"fmt"
	"sync"
	"time"
)

/**
锁
写锁
获取到写锁之后，获取写锁看看会不会阻塞
结论会阻塞
*/

func main() {
	var rw sync.RWMutex

	fmt.Println("准备获取写锁")
	rw.Lock()
	fmt.Println("已经获取到写锁")

	//这里要异步释放写锁，不然编译会提醒死锁
	go func(rw *sync.RWMutex) {
		time.Sleep(time.Second * 5)
		fmt.Println("释放写锁")
		rw.Unlock()
	}(&rw)

	fmt.Println("准备获取读锁")
	rw.RLock()
	fmt.Println("已经获取到读锁")
}
