package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")

	reader := bufio.NewReader(file)

	// 每次读取固定字节
	//ctx := make([]byte, 10)
	//n, _ := reader.Read(ctx)
	//fmt.Println(string(ctx[:n]))

	// 每次读取一个字节
	//ctx, _ := reader.ReadByte()
	//fmt.Println(string(ctx))

	// 读取到某个字节时停止，返回字节切片
	//ctx, _ := reader.ReadBytes('\n')
	//fmt.Println(string(ctx))

	// 读取到某个字节时停止，返回字符串
	//str, _ := reader.ReadString('\n')
	//fmt.Println(str)

	// 每次读取一行
	//ctx, isPrefix, _ := reader.ReadLine()
	//fmt.Println(string(ctx), isPrefix)

	// 重定向到输出对象
	reader.WriteTo(os.Stdout)

}
