package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	string1 := 1
	string2 := "string2"
	string3 := "string3"
	string4 := "string4"
	string5 := 5

	fmt.Fprintf(file, "%d, %s, %s\n", string1, string2, string3)
	fmt.Fprintln(file, string4)
	fmt.Fprint(file, string5)
}
