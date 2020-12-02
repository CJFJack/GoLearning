package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var locker sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("A:lock before")

		locker.RLock()
		fmt.Println("A:locked")
		time.Sleep(time.Second * 5)

		locker.RUnlock()
		fmt.Println("A:Unlock")

		wg.Done()
	}()

	go func() {
		fmt.Println("B:lock before")

		locker.RLock()
		fmt.Println("B:locked")
		time.Sleep(time.Second * 5)

		locker.RUnlock()
		fmt.Println("B:Unlock")

		wg.Done()
	}()

	wg.Wait()
}
