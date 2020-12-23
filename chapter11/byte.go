package main

import "fmt"

func main() {
	b := []byte("ab")
	fmt.Printf("%T, %#v\n", b, b)
}
