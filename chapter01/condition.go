package main

import "fmt"

func main() {
	var y string
	fmt.Println("please input y:")
	_, _ = fmt.Scan(&y)
	if y == "yes" {
		fmt.Print("y = " + y)
	} else {
		fmt.Println("y != yes")
	}

	var score int
	_, _ = fmt.Scan(&score)
	if score > 90 {
		println("A")
	} else if score > 80 {
		println("B")
	} else {
		println("C")
	}

	var score2 int
	_, _ = fmt.Scan(&score2)
	switch {
	case score2 > 90:
		println("A")
	case score2 > 80:
		println("B")
	default:
		println("C")
	}
}
