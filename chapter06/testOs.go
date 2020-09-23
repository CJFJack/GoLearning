package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "www.baidu.com")

	// 执行完后输出结果
	//bytes, err := cmd.Output()
	//fmt.Println(string(bytes), err)

	// 流管道接收实时输出
	output, err := cmd.StdoutPipe()
	cmd.Start()
	fmt.Println(err)
	io.Copy(os.Stdout, output)
	cmd.Wait()

}
