package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Stat("password.txt")
	if err != nil {
		res := os.IsNotExist(err)
		if res {
			fmt.Println("文件不存在")
		} else {
			fmt.Println("文件存在")
		}
	}
	fmt.Println(file)

}
