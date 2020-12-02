package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	select {
	case e := <-ch1:
		fmt.Println(e)
	case e := <-ch2:
		fmt.Println(e)
	}
}
