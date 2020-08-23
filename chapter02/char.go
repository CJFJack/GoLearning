package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	uni := "abc我爱中华人民共和国"
	fmt.Println(len(uni))
	fmt.Println(len([]rune(uni)))
	fmt.Println(utf8.RuneCountInString(uni))

	fmt.Println(string([]rune(uni)))
	fmt.Println(string([]byte(uni)))
}
