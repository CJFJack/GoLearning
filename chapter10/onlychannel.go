package main

import "fmt"

func main() {
	channel := make(chan int, 10)

	func(channel chan<- int) {
		channel <- 1
	}(channel)

	func(channel <-chan int) {
		fmt.Println(<-channel)
	}(channel)

}
