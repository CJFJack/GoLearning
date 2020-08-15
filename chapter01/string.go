package main

import "fmt"

func main() {
	var name string = "Jack"
	fmt.Println(name)

	// 操作
	fmt.Println("My name is " + name)
	name += " hahaha"
	fmt.Println(name + "xxx")
}
