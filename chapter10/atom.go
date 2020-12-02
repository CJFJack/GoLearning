package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	var wg sync.WaitGroup
	count := 5
	ceil := 10000

	for i := 0; i < count; i++ {
		wg.Add(2)
		go func() {
			for i := 0; i < ceil; i++ {
				//counter++
				atomic.AddInt64(&counter, 1)
				//time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < ceil; i++ {
				//counter--
				atomic.AddInt64(&counter, -1)
				//time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
		wg.Wait()
	}

	fmt.Println(counter)

}
