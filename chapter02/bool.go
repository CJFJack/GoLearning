package main

import "fmt"

func main() {
	var zero bool
	isBoy := true
	isGirl := false
	fmt.Println(zero, isBoy, isGirl)
	fmt.Printf("%T, %t\n", isBoy, isGirl)

	// 逻辑运算（与 && 或 || 非 !）
	fmt.Println("逻辑运算")
	fmt.Println(true && false)
	fmt.Println(true && true)
	fmt.Println(false && false)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(true || true)
	fmt.Println(!true)
	fmt.Println(!false)

	// 关系运算（等于 ==  不等于 !=）
	fmt.Println("关系运算")
	fmt.Println(isGirl == isBoy)
	fmt.Println(isGirl != isBoy)



}
