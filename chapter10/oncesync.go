package main

import (
	"fmt"
	"sync"
)

func hi() {
	fmt.Println("hi")
}

func main() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		once.Do(hi)
	}
}
