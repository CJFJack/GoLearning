package main

import "fmt"

func main() {
	name := []int{1, 2, 3, 4, 5}
	//subName := name[2:3]
	subName := name[1:2:3]
	fmt.Println(len(subName), cap(subName))
	fmt.Println(name, subName)
	subName = append(subName, 6)
	fmt.Println(name, subName)
}
