package main

import (
	"fmt"
)

func main() {

	channel := make(chan struct{})

	go func() {
		channel <- struct{}{}
		close(channel)
	}()

	ctx, close := <-channel
	fmt.Println(ctx)
	fmt.Println(close)

}
