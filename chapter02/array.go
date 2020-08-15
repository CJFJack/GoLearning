package main

import "fmt"

func main() {
	var a int
	fmt.Printf("%#v\n", a)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%v\n", a)

	b := [5]string{1: "a"}
	fmt.Printf("%#v\n", b)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%v\n", b)

	var c [3][2]int
	fmt.Printf("%#v", c)
}
