package main

import "fmt"

func main() {
	letters := "abcdefghijkl"
	for i := 0; i < len(letters); i++ {
		fmt.Printf("%c\n", letters[i])
	}

	letters2 := "我爱中华人民共和国"
	for _, v := range letters2 {
		fmt.Printf("%q\n", v)
	}

	var (
		index = 0
		sum   = 0
	)
	for {
		sum += index
		index++
		if index > 100 {
			break
		}
	}
	println(sum)

}
