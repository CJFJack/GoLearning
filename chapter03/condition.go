package main

import "fmt"

func main() {
	var y string
	fmt.Println("please input y:")
	_, _ = fmt.Scan(&y)
	if y == "yes" {
		fmt.Print("y = " + y)
	}
}
