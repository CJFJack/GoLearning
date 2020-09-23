package main

import (
	"flag"
	"fmt"
)

func main() {
	var host string
	var port int
	var h bool
	var help bool

	flag.StringVar(&host, "H", "127.0.0.1", "连接IP")
	flag.IntVar(&port, "P", 22, "连接端口")
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&help, "help", false, "帮助")

	// 解释参数
	flag.Parse()

	if h || help {
		fmt.Println("usage: testflag.exe [-H 127.0.0.1] [-P 3306]")
		flag.PrintDefaults()
		return
	}

	fmt.Println(host, port, h, help)
	// 打印其他参数，返回切片
	fmt.Println(flag.Args())

}
