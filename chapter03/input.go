package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := ""
	fmt.Println("Please input your name:")
	_, _ = fmt.Scan(&name)
	fmt.Println("Your name is " + name)

	var age int
	fmt.Println("Please input your age:")
	_, _ = fmt.Scan(&age)
	fmt.Println("Your age is " + strconv.Itoa(age))

	var msg string
	fmt.Println("Please input message:")
	_, _ = fmt.Scan(&msg)
	fmt.Println("The message is " + msg)
}
