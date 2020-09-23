package main

import (
	"log"
	"os"
)

func main() {
	// 方式一：

	// 设置格式
	log.SetFlags(log.Flags())
	// 设置前缀
	log.SetPrefix("main: ")

	// 正常日志
	log.Println("第一条Println日志")
	// 不可恢复的错误日志
	//log.Fatalln("Fatalln日志")
	// 可恢复的错误日志
	//log.Panicln("Panic日志")

	// 方式二：
	// New参数 输出（控制台、文件）、日志前缀、日志格式
	logger := log.New(os.Stdout, "logger: ", log.Flags())
	//logger.SetPrefix("main: ")
	logger.Println("logger Println日志")

}
