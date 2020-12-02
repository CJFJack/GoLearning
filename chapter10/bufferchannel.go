package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	channel := make(chan int, 3)
	go func() {
		for i := 0; i < 10; i++ {
			//time.Sleep(time.Second * 2)
			//fmt.Println("<-", i)
			channel <- i
		}
		close(channel)
	}()

	wg.Add(1)
	go func() {
		for c := range channel {
			//time.Sleep(time.Second * 2)
			fmt.Println(c)
		}
		wg.Done()
	}()

	wg.Wait()

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("<-", i)
		}
		wg.Done()
	}()

	wg.Wait()

}
