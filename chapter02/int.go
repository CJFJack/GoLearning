package main

import "fmt"

func main() {
	var age int = 29
	fmt.Printf("%T, %d\n", age, age)

	// 操作
	fmt.Println(age * 3)
	fmt.Println(age / 3)
	fmt.Println(age % 3)
	age++
	fmt.Println(age)
	age--
	fmt.Println(age)

	fmt.Println(age > 30)

	// 位运算
	// 与 & 转换成二进制再在相同位数上与，再转换回十进制
	fmt.Println(7 & 2)
	// 或 | 转换成二进制再在相同位数上与或，再转换回十进制
	fmt.Println(7 ^ 2)
	fmt.Println(7 | 2)
	// 左移 <<
	fmt.Println(2 << 1)

	var a byte = 'A'
	var w rune = '中'
	fmt.Println(a)
	fmt.Println(w)
	fmt.Printf("%T, %c, %b, %x. %o\n", a, a, a, a, a)
	fmt.Printf("%T, %U\n", w, w)

	fmt.Printf("%-5d\n", age)
	fmt.Printf("%05d\n", age)

}
