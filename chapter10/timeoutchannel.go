package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	result := make(chan interface{})

	go func() {
		time.After(time.Second * 5)
	}()

	go func() {
		r := rand.Intn(10)
		time.Sleep(time.Second * time.Duration(r))
		result <- r
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout")
	case res := <-result:
		fmt.Println(res)
	}
}
