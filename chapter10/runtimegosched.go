package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(i int) {
			for c := 'A'; c <= 'Z'; c++ {
				fmt.Printf("%d:%c\n", i, c)
				runtime.Gosched()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

}
