package main

import (
	"fmt"
	"os"
)

func main() {
	path, err := os.Getwd()
	fmt.Println(path)
	fmt.Println(err)

	file, err := os.Executable()
	fmt.Println(file)
	fmt.Println(err)

}
