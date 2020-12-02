package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

func main() {
	// 创建文件指针
	file, _ := os.Create("test.csv")
	defer file.Close()

	// 创建缓存对象
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// 创建csv指针对象
	csvWriter := csv.NewWriter(writer)
	// 写入csv文件
	csvWriter.Write([]string{"Id", "Name", "StartTime", "EndTime", "User"})
}
