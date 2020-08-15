package main

import "fmt"

func main() {
	var f float64 = 1.75
	fmt.Printf("%T, %f\n", f, f)
	fmt.Printf("%T, %.2f\n", f, f)
	fmt.Printf("%T, %5.2f\n", f, f)

	// 操作
	// 算数运算
	fmt.Println(1.1 + 1.2)
}
