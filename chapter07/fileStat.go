package main

import (
	"fmt"
	"os"
)

func fileIsExits(err error) bool {
	if err == nil {
		return true
	} else {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
}

func main() {
	fileInfo, err := os.Stat("password.txt")
	isExists := fileIsExits(err)
	if !isExists {
		fmt.Println("文件不存在")
		return
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Mode())
}
