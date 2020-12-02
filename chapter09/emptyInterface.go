package main

import "fmt"

type EmptyInf interface{}

func Print(a ...interface{}) {
	fmt.Println(a)
}

func main() {
	var emptyInf EmptyInf
	emptyInf = 1
	emptyInf = "a"
	emptyInf = true
	fmt.Println(emptyInf)

	var emptyInf2 interface{}
	fmt.Println(emptyInf2)

	Print("a", 1)
}
