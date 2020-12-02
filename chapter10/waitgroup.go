package main

import (
	"fmt"
	"sync"
	"time"
)

func mainChars(prefix string) {
	//fmt.Println(prefix)
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond)
	}
}

func chars(prefix string, wg *sync.WaitGroup) {
	//fmt.Println(prefix)
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond)
	}
	wg.Done()
}

func main() {
	// go调用
	var wg sync.WaitGroup
	wg.Add(1)
	go chars("goroutine", &wg)

	// 主例程
	mainChars("main")

	fmt.Println("wait")
	wg.Wait()
	fmt.Println("over")

}
