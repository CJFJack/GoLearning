package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	lines  int
	follow string
	h      bool
	help   bool
)

func main() {

	flag.IntVar(&lines, "n", 10, "文件最后n行信息")
	flag.StringVar(&follow, "f", "", "跟踪文件输入")
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&help, "help", false, "帮助")

	// 解释参数
	flag.Parse()

	if h || help {
		fmt.Println("usage: tail [-n 10] [-f filename]")
		flag.Usage()
		return
		//os.Exit(0)
	}

	file, _ := os.Open(follow)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		lineText, _ := reader.ReadString('\n')
		fmt.Print(lineText)
	}

}
