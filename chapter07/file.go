package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 打开文件
	file, err := os.Open("password")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 延迟关闭文件
	defer file.Close()

	// 读取文件
	ctx := make([]byte, 3)
	file.Seek(3, os.SEEK_SET)
	for {
		n, readErr := file.Read(ctx)
		if readErr == io.EOF {
			break
		}
		fmt.Print(string(ctx[:n]))
	}

	// 写入文件
	writeFile, writeErr := os.Create("name.txt")
	if writeErr != nil {
		return
	}
	writeFile.Write([]byte("jack\n"))
	writeFile.Write([]byte("love\n"))
	writeFile.Write([]byte("lulu\n"))

	// 追加文件
	appendFile, appendErr := os.OpenFile("name.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if appendErr != nil {
		return
	}
	appendFile.Write([]byte("hahaha"))

}
