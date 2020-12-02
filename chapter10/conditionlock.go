package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var locker sync.Mutex
	cond := sync.NewCond(&locker)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		// 获取锁
		cond.L.Lock()
		// 延迟释放锁
		defer cond.L.Unlock()
		fmt.Println("condition wait before")
		// 等待锁
		cond.Wait()
		fmt.Println("condition wait after")

		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 5)
		// 广播
		fmt.Println("broadcast")
		cond.Broadcast()

		wg.Done()
	}()

	wg.Wait()
}
