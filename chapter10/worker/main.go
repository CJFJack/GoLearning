package main

import (
	"fmt"
	"learning/chapter10/worker/pool"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var wgAddTask sync.WaitGroup

	worker := pool.NewPool(2)

	// 添加任务
	wgAddTask.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			worker.AddTask(func() interface{} {
				return i
			})
			time.Sleep(time.Millisecond)
		}
		wgAddTask.Done()
	}()

	// 执行任务
	worker.Start()

	// 获取结果
	wg.Add(1)
	go func() {
		for r := range worker.Results {
			fmt.Println(r)
		}
		wg.Done()
	}()

	// 等待添加任务的例程结束
	wgAddTask.Wait()

	// 关闭任务池，等待任务执行完成
	worker.Close()

	// 再等待获取结果的例程结束
	wg.Wait()
}
