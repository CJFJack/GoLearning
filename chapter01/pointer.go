package main

import "fmt"

func main() {
	var a = 8
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	// 通过指针地址修改值
	*b = 20
	fmt.Println(a)

	// 空指针
	var c *int
	fmt.Println(c)
}
