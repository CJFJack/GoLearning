package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f float64 = 3.14
	fmt.Println(strconv.FormatFloat(f, 'f', 2, 64))
	fmt.Printf("%.3f\n\n", f)

	var s = "10"
	i, err := strconv.Atoi(s)
	fmt.Println(i, err)

	var str = "6.258"
	r, err := strconv.ParseFloat(str, 64)
	fmt.Println(r, err)
}
