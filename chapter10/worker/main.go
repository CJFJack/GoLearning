package main

import (
	"fmt"
	"learning/chapter10/worker/pool"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	worker := pool.NewPool(5)
	worker.AddTask(func() interface{} {
		return 1
	})
	worker.AddTask(func() interface{} {
		return 2
	})
	worker.AddTask(func() interface{} {
		return 3
	})

	worker.Start()

	// 获取结果
	wg.Add(1)
	go func() {
		for r := range worker.Results {
			fmt.Println(r)
		}
		wg.Done()
	}()

	// 先关闭任务池
	worker.Close()

	// 再关闭获取结果的例程
	wg.Wait()
}
