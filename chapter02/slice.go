package main

import "fmt"

func main() {
	var slice []string

	slice = append(slice, "1")

	//fmt.Println(slice)

	names := []string{"1", "2", "3", "4", "5"}
	//fmt.Println(names)
	//fmt.Println(len(names))
	//fmt.Println(cap(names))
	fmt.Println(names[3:])
	fmt.Println(names[4:])
	copy(names[4:], names[3:])
	fmt.Println(names)
}
