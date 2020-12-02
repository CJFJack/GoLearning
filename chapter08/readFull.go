package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	// 读取并保存整个文件的内容到字节切片或者内存中
	// 方法一
	file, _ := os.Open("Reader.go")
	ctx := make([]byte, 10)
	buffer := make([]byte, 0, 1024*1024)

	defer file.Close()

	for {
		n, err := file.Read(ctx)
		if err != nil {
			break
		}
		buffer = append(buffer, ctx[:n]...)
	}

	fmt.Println(string(buffer))

	// 方法二
	file2, _ := os.Open("filepath.go")
	builder := new(strings.Builder)
	io.Copy(builder, file2)

	defer file2.Close()

	fmt.Println(builder.String())

	// 方法三
	file3, _ := os.Open("copyFile.go")
	defer file3.Close()
	newBuffer := bytes.NewBuffer([]byte(""))
	io.Copy(newBuffer, file3)
	fmt.Println(newBuffer.String())

	// 方法四
	file4, err := ioutil.ReadFile("stringsBuilder.go")
	fmt.Println(string(file4), err)

}
