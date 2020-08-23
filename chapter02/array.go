package main

import "fmt"

func main() {
	var a int
	fmt.Printf("%#v\n", a)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%v\n", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b == c)
	c[1] = 10
	fmt.Println(c)
	fmt.Println(c[1:3])
}
