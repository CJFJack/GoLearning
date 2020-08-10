package main

import "fmt"

func main() {
	const (
		C1 = (iota + 1) * 100
		C2
		C3
	)
	fmt.Println(C1, C2, C3)

	a := 1
	{
		fmt.Println(a)
		b := 2
		{
			fmt.Println(b)
			a := 3
			fmt.Println(a)
		}
	}
	fmt.Printf("%T, %d", a, a)
}
