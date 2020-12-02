package main

import (
	"fmt"
	"sync"
)

type Connection struct {
	id   int
	addr string
}

func main() {
	addr := "192.168.1.1"
	index := 0

	pool := sync.Pool{
		New: func() interface{} {
			index++
			fmt.Println("New:", index)
			return &Connection{id: index, addr: addr}
		},
	}

	c := pool.Get()
	fmt.Println("Get", c)

	pool.Put(c)
	fmt.Println("Put")

	c = pool.Get()
	fmt.Println("Get", c)

	c = pool.Get()
	fmt.Println("Get", c)
}
