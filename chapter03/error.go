package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer")
		fmt.Println(recover())
	}()
	fmt.Println("before")
	panic("自定义panic")
	fmt.Println("after")
}
