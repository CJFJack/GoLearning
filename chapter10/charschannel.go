package main

import (
	"fmt"
)

func main() {
	const count = 3
	channel := make(chan int)
	for i := 0; i < count; i++ {
		go func(prefix int) {
			for ch := 'A'; ch < 'Z'; ch++ {
				fmt.Printf("%d: %c\n", prefix, ch)
			}
			channel <- prefix
		}(i)
	}

	// 从管道读取count次，证明count个工作例程已结束
	for i := 0; i < count; i++ {
		fmt.Println("over: ", <-channel)
	}

}
