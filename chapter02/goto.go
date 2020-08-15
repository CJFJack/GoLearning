package main

import (
	"fmt"
)

func main() {
	var (
		index = 1
		total = 1
	)
	const max = 100
START:
	index++
	total += index
	if index <= max {
		goto START
	}
	fmt.Println(total)
}
