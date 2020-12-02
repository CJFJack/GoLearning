package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//wg.Add(10)
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		//wg.Wait()
	}
	wg.Wait()

}
