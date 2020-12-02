package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义
	var channel chan int
	fmt.Printf("%T, %#v\n", channel, channel)

	// 初始化
	channel = make(chan int)
	fmt.Printf("%T, %#v\n", channel, channel)

	// 写，必须要有其他例程进行读操作
	go func() {
		//time.Sleep(time.Second * 5)
		channel <- 1
	}()

	// 读，必须要有其他例程进行写操作
	time.Sleep(time.Second * 5)
	name := <-channel
	fmt.Println(name)
	time.Sleep(time.Second)
}
