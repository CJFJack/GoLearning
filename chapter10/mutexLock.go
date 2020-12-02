package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	const count = 1000
	var (
		a      = 10000
		b      = 10000
		wg     sync.WaitGroup
		locker sync.Mutex
	)

	wg.Add(2)

	// a-->b转账100次
	go func() {
		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if a > money {
				locker.Lock()
				a -= money
				//time.Sleep(time.Microsecond)
				b += money
				locker.Unlock()
			}
		}
		wg.Done()
	}()

	// b-->a转账100次
	go func() {
		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if b > money {
				locker.Lock()
				b -= money
				//time.Sleep(time.Microsecond)
				a += money
				locker.Unlock()
			}
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("a: %d, b: %d, sum: %d", a, b, a+b)

}
