package main

import "fmt"

func main() {
	i := 1
	j := 1
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d  ", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
